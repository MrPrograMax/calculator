package service

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"testing"
	"time"

	"calculator/internal/pkg/model"
)

const benchJSON = `[
  { "type": "calc", "op": "+", "var": "x",       "left": 10, "right": 2   },
  { "type": "calc", "op": "*", "var": "y",       "left": "x", "right": 5   },
  { "type": "calc", "op": "-", "var": "q",       "left": "y", "right": 20  },
  { "type": "calc", "op": "+", "var": "unusedA", "left": "y", "right": 100 },
  { "type": "calc", "op": "*", "var": "unusedB", "left": "unusedA", "right": 2 },
  { "type": "print",           "var": "q"                     },
  { "type": "calc", "op": "-", "var": "z",       "left": "x", "right": 15  },
  { "type": "print",           "var": "z"                     },
  { "type": "calc", "op": "+", "var": "ignoreC", "left": "z", "right": "y" },
  { "type": "print",           "var": "x"                     }
]`

func BenchmarkExec_SampleWorkflow(b *testing.B) {
	var expressions []*model.Expression
	if err := json.Unmarshal([]byte(benchJSON), &expressions); err != nil {
		b.Fatalf("unmarshal benchJSON: %v", err)
	}
	svc := NewService(&logrus.Logger{})
	ctx := context.Background()

	b.ResetTimer()

	start := time.Now()
	for i := 0; i < b.N; i++ {
		if _, err := svc.Exec(ctx, expressions); err != nil {
			b.Fatalf("Exec failed: %v", err)
		}
	}
	elapsed := time.Since(start)

	secPerOp := elapsed.Seconds() / float64(b.N)
	b.ReportMetric(secPerOp, "sec/op")
}
