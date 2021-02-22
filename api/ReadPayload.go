package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const DEFAULTREADTIMEOUT = time.Second

// ReadPayload reads the payload that the scanner read
func (a *API) ReadPayload(ctx echo.Context) error {
	fmt.Println("Read")
	payload, err := scanner.Read()
	if err != nil || (payload != nil && len(payload) < 1) {
		if err == nil {
			err = fmt.Errorf("unknown read error")
		}
		return SendError(ctx, http.StatusNotFound, err.Error())
	}
	var read Read = Read{
		Payload: new(string),
	}
	*read.Payload = string(payload)
	fmt.Println(*read.Payload)
	err = ctx.JSON(http.StatusOK, read)
	return err
}
