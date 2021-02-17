// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /disable_code)
	DisableCode(ctx echo.Context) error

	// (POST /enable_code)
	EnableCode(ctx echo.Context) error

	// (GET /info)
	GetInfo(ctx echo.Context) error

	// (POST /light)
	Light(ctx echo.Context) error

	// (GET /read)
	ReadPayload(ctx echo.Context) error

	// (GET /test)
	GetTest(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// DisableCode converts echo context to params.
func (w *ServerInterfaceWrapper) DisableCode(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DisableCode(ctx)
	return err
}

// EnableCode converts echo context to params.
func (w *ServerInterfaceWrapper) EnableCode(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.EnableCode(ctx)
	return err
}

// GetInfo converts echo context to params.
func (w *ServerInterfaceWrapper) GetInfo(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetInfo(ctx)
	return err
}

// Light converts echo context to params.
func (w *ServerInterfaceWrapper) Light(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Light(ctx)
	return err
}

// ReadPayload converts echo context to params.
func (w *ServerInterfaceWrapper) ReadPayload(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ReadPayload(ctx)
	return err
}

// GetTest converts echo context to params.
func (w *ServerInterfaceWrapper) GetTest(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTest(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/disable_code", wrapper.DisableCode)
	router.POST(baseURL+"/enable_code", wrapper.EnableCode)
	router.GET(baseURL+"/info", wrapper.GetInfo)
	router.POST(baseURL+"/light", wrapper.Light)
	router.GET(baseURL+"/read", wrapper.ReadPayload)
	router.GET(baseURL+"/test", wrapper.GetTest)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RWX3PjNBD/KpqFRxM7CVx7frsezE0HGI6Wt04mo1qbWDeOpJPWLZkbf3dGku04tRMK",
	"LQO82drVb//9dldfoNA7oxUqcpB/AYvOaOUw/DjiVLvuyJ8UWhEq8p/cmEoWnKRW6SenVbhQlLjj/utr",
	"ixvI4av0AJ9GqUsjLDRNk4BAV1hpPArkrUG2Q+f4FtmjpJK1Z4UWyLgSnRD87RYxeiZwTXsT3DxGRVXv",
	"IhZ/4LLi9xVGOK/uIAkKkN8BcjVf+n+uLiGBzxYS4FUFqwTwd74zFXq0VinaAkdWqi00CVjkwhs3Vhu0",
	"JGMODd9XOgqOvaISWStkVHJi/sAVXCm0LGANrb69uMwW2Xx58fbNYmy86U/0/ScsCEJuAtRaqo0eu1Vy",
	"Kx65xfUDWidj9Y796zRYpzF0Zz7Lsqkc7LTAaoxlrBZ1QSyKh0Affn7z3RSQ0xsK/glOEyXtxCyIh4CL",
	"bJFly2x+FvRk0D3uiaDnz8x9ZPgo6551Y6O3B4ZDAhttd5wgB6loOai1VIRbtCHLbQucQurkU75a/FxL",
	"i8LzvbXYqa+mInmUVJRrbehcX1VyWxLTQXDUUMo3GyQgpGu/HIkn/dTpjH11WNRW0v7Wt3nM4BV3snhX",
	"U9mPG3/n3p8eIEoiE8dLR/5jx6+49aGnv958897Pgdu26zwZ2buP1x5J0pCfPWEgm81nmc+MNqi4kZDD",
	"cpbNPDEMpzI4mbbhrrt6G+1C/jwXwsS8FpDD91HrfayCLww6utJi/5cm7Zhh/Rg8N4QPihMMHk/mNiK2",
	"lQ+oDtMzzr3Bylhk2SnLvV76ZK80CXz7d64FN9PInz9J9Q/qf5PpGM9/MNFdL21xYhbcINVWMb80wpZ+",
	"Okgd05vhhoPkSYU+IF17A9Nhvs67Y7gRJzJ/WxcFOrepq2rPbAgIBYvaL8pcGI+nyflTEL8WL10sz9lE",
	"HIb6syjpkLoRv+kL+O8TsntwnSGkYzy8pdrn44bQMu1p6tg9omqDESM23iAXH9uX2z/IyBDBs5nYB/LC",
	"vBFGFp7N2y8/TrXob/7qaxQ+utIteMjvjlb73apZean10yMIa1u1iz1P0x1XW51fZhcZNKvmjwAAAP//",
	"+4Mij8MMAAA=",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
