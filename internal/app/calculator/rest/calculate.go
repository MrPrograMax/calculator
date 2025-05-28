package rest

import (
	myerror "calculator/internal/error"
	"calculator/internal/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Calculate обрабатывает список инструкций calc/print и возвращает результаты команд print.
// @Summary      Выполнить инструкции калькулятора
// @Description  Обрабатывает список инструкций:
// @Description    • calc – вычисляет арифметическую операцию (с эмуляцией задержки 50 ms)
// @Description    • print – возвращает значение переменной
// @Tags         Calculator
// @Accept       json
// @Produce      json
// @Param        expressions  body      []model.Expression  true  "Список инструкций calc/print"
// @Success      200          {array}   model.Result       "Результаты print в порядке вызова"
// @Failure      400          {object}  map[string]string  "Ошибка валидации или выполнения"
// @Failure      500          {object}  map[string]string  "Внутренняя ошибка сервера"
// @Router       /api/calculate [post]
func (h *Handler) Calculate(c *gin.Context) {
	var input []*model.Expression

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
	}

	response, err := h.service.calculatorREST.Exec(c.Request.Context(), input)
	if err != nil {
		h.service.logger.Errorf("[app][calculator][rest] Exec: %v", err.Error())
		newErrorResponse(c, http.StatusInternalServerError, myerror.ErrSomethingWentWrong.Error())
		return
	}

	c.JSON(http.StatusOK, getCalculateResponse{
		Items: response,
	})
}
