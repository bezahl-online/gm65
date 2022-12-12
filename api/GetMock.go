package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *API) GetMock(ctx echo.Context, params GetMockParams) error {
	scanner.SetCode(params.Code)
	if err := SendStatus(ctx, http.StatusOK, "OK"); err != nil {
		return err
	}
	return nil
}
