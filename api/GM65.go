package api

import (
	"fmt"
	"log"
	"net/http"

	"bezahl.online/gm65/device"
	"github.com/labstack/echo/v4"
)

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=types.cfg.yaml gm65.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=server.cfg.yaml gm65.yaml

// API is the api interface type
type API struct{}

var e *echo.Echo = echo.New()
var scanner device.Scanner

func init() {
	testServer := &API{}

	scanner = device.Scanner{
		Config: device.Config{
			SerialPort: "/dev/ttyACM0",
			Baud:       9600,
		},
	}

	err := scanner.Open()
	if err != nil {
		log.Fatal(err)
	}

	RegisterHandlers(e, testServer)
}

// GetTest returns status ok
func (a *API) GetTest(ctx echo.Context) error {
	var err error
	err = SendStatus(ctx, http.StatusOK, "OK")
	if err != nil {
		return err
	}
	return nil
}

// SendStatus function wraps sending of an error in the Error format, and
// handling the failure to marshal that.
func SendStatus(ctx echo.Context, code int, message string) error {
	statusError := Status{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, statusError)
	return err
}

// SendError function wraps sending of an error in the Error format
// returns sent error message or new error if sending fails
func SendError(ctx echo.Context, code int, message string) error {
	status := Status{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, status)
	if err != nil {
		return err
	}
	return fmt.Errorf(status.Message)
}