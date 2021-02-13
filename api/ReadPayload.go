package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ReadPayload reads the payload that the scanner read
func (a *API) ReadPayload(ctx echo.Context) error {
	payload, err := scanner.Read()
	if err != nil {
		return SendError(ctx, http.StatusBadRequest, err.Error())
	}
	var read Read = Read{
		Payload: new(string),
	}
	*read.Payload=string(payload)
	err = ctx.JSON(http.StatusOK,read)
	return err
}
