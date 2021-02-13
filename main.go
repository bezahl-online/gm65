package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"bezahl.online/gm65/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	swagger, err :=api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	serverPort, err := strconv.Atoi(strings.Split(swagger.Servers[0].URL, ":")[2])
	port := &serverPort
	e := echo.New()
	e.Use(middleware.CORS())
	server := &api.API{}

	api.RegisterHandlers(e, server)

	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))

}
