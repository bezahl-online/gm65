package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"bezahl.online/gm65/api"
	"bezahl.online/gm65/param"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}
	if *param.Disconnected {
		fmt.Println("***\n*** test mode - not reporting missing physical device ***\n***")
	}
	serverPort, err := strconv.Atoi(strings.Split(swagger.Servers[0].URL, ":")[2])
	if err != nil {
		log.Fatal(err)
	}
	port := &serverPort
	e := echo.New()
	e.Use(middleware.CORS())
	server := &api.API{}

	api.RegisterHandlers(e, server)

	if len(os.Getenv("PRODUCTIVE")) > 0 {
		e.Logger.Fatal(e.StartTLS(fmt.Sprintf("0.0.0.0:%d", *port), "localhost.crt", "localhost.key"))
	} else {
		e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))
	}
}
