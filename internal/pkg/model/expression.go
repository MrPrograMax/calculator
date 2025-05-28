package model

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Expression математическое выражение.
// swagger:model Expression
type Expression struct {
	// Тип инструкции: calc – вычисление, print – вывод результата
	// Enum: [calc print]
	// required: true
	Type Type `json:"type" example:"calc"`

	// Тип вычислительной операции: +, -, *, /
	// example: "+"
	Op *Operation `json:"op,omitempty" example:"+"`

	// Название переменной, в которую записывается результат или печатаем
	// required: true
	// example: x
	Var string `json:"var" example:"x"`

	// Левый операнд: либо число, либо имя переменной
	// oneOf: [integer string]
	// example: 10
	Left RawOperand `json:"left"`

	// Правый операнд: либо число, либо имя переменной
	// oneOf: [integer string]
	// example: 20
	Right RawOperand `json:"right"`
}

// RawOperand умеет распознавать JSON-число или JSON-строку
// swagger:model RawOperand
type RawOperand struct {
	// internal flag, не показывать в Swagger
	IsVar bool `json:"-"`

	// internal value, не показывать в Swagger
	IntVal int64 `json:"-"`

	// internal name, не показывать в Swagger
	VarName string `json:"-"`
}

func (r *RawOperand) UnmarshalJSON(b []byte) error {
	// Сначала пробуем чистый JSON-числовой литерал.
	if err := json.Unmarshal(b, &r.IntVal); err == nil {
		r.IsVar = false
		return nil
	}

	// Кейс когда цифра является строкой, например "10".
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		if iv, err2 := strconv.ParseInt(s, 10, 64); err2 == nil {
			r.IntVal = iv
			r.IsVar = false
			return nil
		}

		r.VarName = s
		r.IsVar = true
		return nil
	}
	return fmt.Errorf("invalid operand %s", string(b))
}

// Type Тип инструкции.
// swagger:enum
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
// swagger:enum
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
