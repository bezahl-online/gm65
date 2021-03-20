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
	var read Read = Read{
		Payload: new(string),
	}
	*read.Payload = scanner.GetLatest()
	fmt.Println(*read.Payload)
	return ctx.JSON(http.StatusOK, read)
}
