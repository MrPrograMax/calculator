package model

// Result результат вычисления
// swagger:model Result
type Result struct {
	// Название переменной, результат которой выводится
	// required: true
	// example: x
	Var string `json:"var"`

	// Значение переменной
	// required: true
	// example: 42
	Value int64 `json:"value"`
}
