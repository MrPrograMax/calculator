// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/calculate": {
            "post": {
                "description": "Обрабатывает список инструкций:\n• calc – вычисляет арифметическую операцию (с эмуляцией задержки 50 ms)\n• print – возвращает значение переменной",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Calculator"
                ],
                "summary": "Выполнить инструкции калькулятора",
                "parameters": [
                    {
                        "description": "Список инструкций calc/print",
                        "name": "expressions",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Expression"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Результаты print в порядке вызова",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Result"
                            }
                        }
                    },
                    "400": {
                        "description": "Ошибка валидации или выполнения",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Expression": {
            "type": "object",
            "properties": {
                "left": {
                    "description": "Левый операнд: либо число, либо имя переменной\noneOf: [integer string]\nexample: 10",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.RawOperand"
                        }
                    ]
                },
                "op": {
                    "description": "Тип вычислительной операции: +, -, *, /\nexample: \"+\"",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.Operation"
                        }
                    ],
                    "example": "+"
                },
                "right": {
                    "description": "Правый операнд: либо число, либо имя переменной\noneOf: [integer string]\nexample: 20",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.RawOperand"
                        }
                    ]
                },
                "type": {
                    "description": "Тип инструкции: calc – вычисление, print – вывод результата\nEnum: [calc print]\nrequired: true",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.Type"
                        }
                    ],
                    "example": "calc"
                },
                "var": {
                    "description": "Название переменной, в которую записывается результат или печатаем\nrequired: true\nexample: x",
                    "type": "string",
                    "example": "x"
                }
            }
        },
        "model.Operation": {
            "description": "Арифметическая операция: +, -, *, /",
            "type": "string",
            "enum": [
                "Unknown",
                "+",
                "-",
                "*",
                "/"
            ],
            "x-enum-varnames": [
                "UnknownOperation",
                "Addition",
                "Subtraction",
                "Multiplication",
                "Division"
            ]
        },
        "model.RawOperand": {
            "type": "object"
        },
        "model.Result": {
            "type": "object",
            "properties": {
                "value": {
                    "description": "Значение переменной\nrequired: true\nexample: 42",
                    "type": "integer"
                },
                "var": {
                    "description": "Название переменной, результат которой выводится\nrequired: true\nexample: x",
                    "type": "string"
                }
            }
        },
        "model.Type": {
            "description": "Тип операции: calc - вычисление, print - вывод результата",
            "type": "string",
            "enum": [
                "Unknown",
                "calc",
                "print"
            ],
            "x-enum-varnames": [
                "UnknownType",
                "Calc",
                "Print"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
