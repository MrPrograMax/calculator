package rest

import (
	myerror "calculator/internal/error"
	"calculator/internal/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Calculate godoc
// @Summary Выполнить вычисления
// @Description Принимает список выражений и возвращает результаты вычислений
// @Tags calculator
// @Accept json
// @Produce json
// @Param input body []model.Expression true "Список выражений для вычисления"
// @Success 200 {object} getCalculateResponse "Результаты вычислений"
// @Failure 400 {object} errorResponse "Некорректные входные данные"
// @Failure 500 {object} errorResponse "Внутренняя ошибка сервера"
// @Router /api/calculate [post]
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
