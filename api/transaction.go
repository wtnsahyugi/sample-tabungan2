package api

import (
	"net/http"
	"sample-tabungan2/entity"
	"sample-tabungan2/service/transaction"

	"github.com/labstack/echo/v4"
)

type TransactionHTTPHandler struct {
	svc *transaction.TransactionService
}

func NewTransactionHTTPHandler(svc *transaction.TransactionService) *TransactionHTTPHandler {
	return &TransactionHTTPHandler{svc: svc}
}

func (h *TransactionHTTPHandler) ListByNoRekening(echoCtx echo.Context) error {
	noRekening := echoCtx.Param("no_rekening")
	data, err := h.svc.ListByRekening(echoCtx.Request().Context(), noRekening)
	if err != nil {
		return handleError(echoCtx, err)
	}

	_ = echoCtx.JSON(http.StatusOK, map[string][]entity.Transaction{
		"mutasi": data,
	})
	return nil
}
