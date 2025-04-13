package model

// Result результат вычисления
// @description Результат вычисления выражения
type Result struct {
	Var   string `json:"Var"`
	Value int64  `json:"Value"`
}
