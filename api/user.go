package api

import (
	"fmt"
	"net/http"
	"sample-tabungan2/entity"
	"sample-tabungan2/service/user"

	"github.com/labstack/echo/v4"
)

type UserHTTPHandler struct {
	svc *user.UserService
}

func NewUserHTTPHandler(svc *user.UserService) *UserHTTPHandler {
	return &UserHTTPHandler{svc: svc}
}

func (h *UserHTTPHandler) Create(echoCtx echo.Context) error {
	req := user.CreateRequest{}
	if err := echoCtx.Bind(&req); err != nil {
		return handleError(echoCtx, entity.ErrInvalidRequestPayload)
	}

	accountNumber, err := h.svc.Create(echoCtx.Request().Context(), req)
	if err != nil {
		return handleError(echoCtx, err)
	}

	_ = echoCtx.JSON(http.StatusCreated, map[string]string{
		"no_rekening": accountNumber,
	})
	return nil
}

func (h *UserHTTPHandler) Deposit(echoCtx echo.Context) error {
	req := user.DepositRequest{}
	if err := echoCtx.Bind(&req); err != nil {
		return handleError(echoCtx, entity.ErrInvalidRequestPayload)
	}

	saldo, err := h.svc.Deposit(echoCtx.Request().Context(), req)
	if err != nil {
		return handleError(echoCtx, err)
	}

	_ = echoCtx.JSON(http.StatusOK, map[string]string{
		"saldo": fmt.Sprintf("%v", saldo),
	})
	return nil
}

func (h *UserHTTPHandler) Withdraw(echoCtx echo.Context) error {
	req := user.WithdrawRequest{}
	if err := echoCtx.Bind(&req); err != nil {
		return handleError(echoCtx, entity.ErrInvalidRequestPayload)
	}

	saldo, err := h.svc.Withdraw(echoCtx.Request().Context(), req)
	if err != nil {
		return handleError(echoCtx, err)
	}

	_ = echoCtx.JSON(http.StatusOK, map[string]string{
		"saldo": fmt.Sprintf("%v", saldo),
	})
	return nil
}
