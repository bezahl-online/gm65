// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

const (
	BasicAuthScopes = "BasicAuth.Scopes"
)

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

// Statusresponse defines model for statusresponse.
type Statusresponse Status
