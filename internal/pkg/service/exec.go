package service

import (
	"context"
	"sync"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/multierr"

	myerror "calculator/internal/error"
	"calculator/internal/pkg/model"
)

func (s Service) Exec(_ context.Context, expressions []*model.Expression) ([]*model.Result, error) {
	type task struct {
		varName    string
		op         model.Operation
		leftVar    string
		rightVar   string
		leftConst  int64
		rightConst int64
	}

	tasks := make(map[string]*task)
	depCount := make(map[string]int)
	dependents := make(map[string][]string)
	results := make(map[string]int64)
	var printOrder []string
	var mu sync.Mutex

	for _, expr := range expressions {
		switch expr.Type {
		case model.Calc:
			if err := validateCalcRequest(expr); err != nil {
				return nil, errors.Wrap(err, "validateCalcRequest")
			}

			t := &task{varName: expr.Var, op: *expr.Op}

			if expr.Left.IsVar {
				t.leftVar = expr.Left.VarName
				depCount[t.varName]++
				dependents[t.leftVar] = append(dependents[t.leftVar], t.varName)
			} else {
				t.leftConst = expr.Left.IntVal
			}

			if expr.Right.IsVar {
				t.rightVar = expr.Right.VarName
				depCount[t.varName]++
				dependents[t.rightVar] = append(dependents[t.rightVar], t.varName)
			} else {
				t.rightConst = expr.Right.IntVal
			}

			tasks[t.varName] = t

		case model.Print:
			printOrder = append(printOrder, expr.Var)

		default:
			return nil, errors.Wrap(myerror.ErrUnknownOperation, string(expr.Type))
		}
	}

	// Канал задач и пул воркеров
	taskCh := make(chan *task, len(tasks))

	var wg sync.WaitGroup

	wg.Add(len(tasks))
	for i := 0; i < len(tasks); i++ {
		go func() {
			for t := range taskCh {
				mu.Lock()
				l := t.leftConst
				if t.leftVar != "" {
					l = results[t.leftVar]
				}
				r := t.rightConst
				if t.rightVar != "" {
					r = results[t.rightVar]
				}
				mu.Unlock()

				time.Sleep(50 * time.Millisecond)

				var res int64
				switch t.op {
				case "+":
					res = l + r
				case "-":
					res = l - r
				case "*":
					res = l * r
				}

				mu.Lock()
				results[t.varName] = res
				for _, dep := range dependents[t.varName] {
					depCount[dep]--
					if depCount[dep] == 0 {
						taskCh <- tasks[dep]
					}
				}
				mu.Unlock()
				wg.Done()
			}
		}()
	}

	// Стартуем задачи без зависимостей
	for name, t := range tasks {
		if depCount[name] == 0 {
			taskCh <- t
		}
	}
	wg.Wait()
	close(taskCh)

	// Собираем результаты в порядке print
	var out []*model.Result
	for _, v := range printOrder {
		if val, ok := results[v]; ok {
			out = append(out, &model.Result{Var: v, Value: val})
		}
	}
	return out, nil
}

func validateCalcRequest(req *model.Expression) error {
	var mErr error

	if req.Op == nil || *req.Op == "" {
		mErr = multierr.Append(mErr, errors.New("empty operation"))
	}

	if req.Type == model.Calc {
		if req.Left.IsVar && req.Left.VarName == "" {
			mErr = multierr.Append(mErr, errors.New("empty left value"))
		}

		if req.Right.IsVar && req.Right.VarName == "" {
			mErr = multierr.Append(mErr, errors.New("empty right value"))
		}
	}

	return mErr
}
