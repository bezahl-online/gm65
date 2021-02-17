package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ReadPayload reads the payload that the scanner read
func (a *API) ReadPayload(ctx echo.Context) error {
	fmt.Println("Read")
	payload, err := scanner.Read()
	if err != nil {
		return SendError(ctx, http.StatusBadRequest, err.Error())
	}
	var read Read = Read{
		Payload: new(string),
	}
	*read.Payload = string(payload)
	fmt.Println(*read.Payload)
	err = ctx.JSON(http.StatusOK, read)
	return err
}
