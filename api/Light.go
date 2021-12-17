package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Light switches the light of the scanner
// on, off or in to standard mode
func (a *API) Light(ctx echo.Context) error {
	var request LightJSONRequestBody
	err := ctx.Bind(&request)
	if err != nil {
		return SendError(ctx, http.StatusBadRequest, err.Error())
	}
	switch *request.Set {
	case SwitchOptEnable:
		err = scanner.LightOn()
		break
	case SwitchOptDisable:
		err = scanner.LightOff()
		break
	case SwitchOptStd:
		err = scanner.LightStd()
		break
	default:
		err = fmt.Errorf("'%s' not implemented", *request.Set)
		if err != nil {
			return SendError(ctx, http.StatusBadRequest, err.Error())
		}
	}
	if err != nil {
		return SendError(ctx, http.StatusGone, err.Error())
	}
	return nil
}
