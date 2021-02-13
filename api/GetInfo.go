package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetInfo returns all data available about the scanner
func (a *API) GetInfo(ctx echo.Context) error {
	model, err := scanner.ProductModel()
	softwareVersion, err := scanner.SoftwareVersion()
	softwareDate, err := scanner.SoftwareDate()
	hardwareVersion, err := scanner.HardwareVersion()

	info := ScannerInfo{
		HardwareVersion: new(string),
		Model:           new(string),
		SoftwareDate:    new(string),
		SoftwareVersion: new(string),
	}
	*info.Model = fmt.Sprintf("GM65 (%d)", int(model))
	*info.HardwareVersion = fmt.Sprintf("%.2f", float32(hardwareVersion)/100)
	*info.SoftwareVersion = fmt.Sprintf("%.2f", float32(softwareVersion)/100)
	*info.SoftwareDate = softwareDate

	err = ctx.JSON(http.StatusOK, info)
	if err != nil {
		return err
	}
	return nil
}
