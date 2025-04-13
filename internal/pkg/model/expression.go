package model

// Expression математическое выражение.
// @Description Математическое выражение для вычисления
type Expression struct {
	Type  Type       `json:"type"`
	Op    *Operation `json:"op,omitempty"`
	Var   string     `json:"var"`
	Left  *string    `json:"left,omitempty"`
	Right *string    `json:"right,omitempty"`
}

// Type Тип инструкции.
// @description Тип операции: calc - вычисление, print - вывод результата
type Type string

const (
	// UnknownType неизвестная операция
	UnknownType Type = "Unknown"
	// Calc вычислить.
	Calc Type = "calc"
	// Print сохранить результат.
	Print Type = "print"
)

// Operation тип операции.
// @description Арифметическая операция: +, -, *, /
type Operation string

const (
	// UnknownOperation неизвестная операция
	UnknownOperation Operation = "Unknown"
	// Addition сложение.
	Addition Operation = "+"
	// Subtraction вычитание.
	Subtraction Operation = "-"
	// Multiplication умножение.
	Multiplication Operation = "*"
	// Division деление.
	Division Operation = "/"
)
