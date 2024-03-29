// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package api

const (
	BasicAuthScopes = "BasicAuth.Scopes"
)

// Defines values for CodeType.
const (
	CodeTypeAll CodeType = "all"

	CodeTypeEan13 CodeType = "ean13"

	CodeTypeEan8 CodeType = "ean8"

	CodeTypeQr CodeType = "qr"

	CodeTypeUpca CodeType = "upca"
)

// Defines values for SwitchOpt.
const (
	SwitchOptDisable SwitchOpt = "disable"

	SwitchOptEnable SwitchOpt = "enable"

	SwitchOptStd SwitchOpt = "std"
)

// enum with available code types
type CodeType string

// Read defines model for read.
type Read struct {
	// the payload that the scanner read
	Payload *string `json:"payload,omitempty"`
}

// ScannerInfo defines model for scanner_info.
type ScannerInfo struct {
	// hardware version
	HardwareVersion *string `json:"hardware_version,omitempty"`

	// product model
	Model *string `json:"model,omitempty"`

	// software date
	SoftwareDate *string `json:"software_date,omitempty"`

	// software version
	SoftwareVersion *string `json:"software_version,omitempty"`
}

// Status defines model for status.
type Status struct {
	// Status code
	Code int32 `json:"code"`

	// Status message
	Message string `json:"message"`
}

// enum with light options
type SwitchOpt string

// Statusresponse defines model for statusresponse.
type Statusresponse Status

// DisableCodeJSONBody defines parameters for DisableCode.
type DisableCodeJSONBody struct {
	// enum with available code types
	CodeType *CodeType `json:"code_type,omitempty"`
}

// EnableCodeJSONBody defines parameters for EnableCode.
type EnableCodeJSONBody struct {
	// enum with available code types
	CodeType *CodeType `json:"code_type,omitempty"`
}

// LightJSONBody defines parameters for Light.
type LightJSONBody struct {
	// enum with light options
	Set *SwitchOpt `json:"set,omitempty"`
}

// GetMockParams defines parameters for GetMock.
type GetMockParams struct {
	// the code to be injected
	Code string `json:"code"`
}

// DisableCodeJSONRequestBody defines body for DisableCode for application/json ContentType.
type DisableCodeJSONRequestBody DisableCodeJSONBody

// EnableCodeJSONRequestBody defines body for EnableCode for application/json ContentType.
type EnableCodeJSONRequestBody EnableCodeJSONBody

// LightJSONRequestBody defines body for Light for application/json ContentType.
type LightJSONRequestBody LightJSONBody
