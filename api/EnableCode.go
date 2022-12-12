package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// EnableCode enables the scanner to read given code type
// on, off or in to standard mode
func (a *API) EnableCode(ctx echo.Context) error {
	var request EnableCodeJSONRequestBody
	err := ctx.Bind(&request)
	if err != nil {
		return SendError(ctx, http.StatusBadRequest, err.Error())
	}
	fmt.Printf("EnableCode %s\n", *request.CodeType)
	switch *request.CodeType {
	case CodeTypeEan13:
		err = scanner.EnableEAN13()
	case CodeTypeEan8:
		err = scanner.EnableEAN8()
	case CodeTypeQr:
		err = scanner.EnableQRCode()
	case CodeTypeUpca:
		err = scanner.EnableUPCa()
	// case CodeTypeAll:
	// 	err = scanner.EnableAll()
	// 	break
	default:
		err = fmt.Errorf("'%s' not implemented", *request.CodeType)
		if err != nil {
			return SendError(ctx, http.StatusBadRequest, err.Error())
		}
	}
	if err != nil {
		return SendError(ctx, http.StatusGone, err.Error())
	}
	return nil
}
