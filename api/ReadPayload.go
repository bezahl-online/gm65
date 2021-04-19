package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const DEFAULTREADTIMEOUT = time.Second

// TODO: debug messages with debug flag
// ReadPayload reads the payload that the scanner read
func (a *API) ReadPayload(ctx echo.Context) error {
	if !scanner.IsConnected() {
		return SendError(ctx, http.StatusNotFound,
			fmt.Sprint(("lost connection to scanner device")))
	}
	var read Read = Read{
		Payload: new(string),
	}
	*read.Payload = scanner.GetCode()
	scanner.ClearCode()
	if len(*read.Payload) > 0 {
		log.Printf("return '%s' to register\n", *read.Payload)
	}
	return ctx.JSON(http.StatusOK, read)
}
