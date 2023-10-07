package api

import (
	"net/http"
	"sample-tabungan2/entity"

	"github.com/labstack/echo/v4"
)

// handleError function to handle error on handler.
func handleError(ctx echo.Context, err error) error {
	code := http.StatusInternalServerError
	switch err {
	case entity.ErrUniqueHpNik,
		entity.ErrInvalidRequestPayload,
		entity.ErrSaldoLessThanWithrawNominal,
		entity.ErrSaldoZero:
		code = http.StatusBadRequest
	}
	_ = ctx.JSON(code, map[string]string{
		"remark": err.Error(),
	})
	return err
}
