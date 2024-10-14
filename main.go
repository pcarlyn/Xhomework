package main

import (
	"homework/web/cmd"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	ltem := e.Group("/")
	cmd.Route(ltem)
	e.Logger.Fatal(e.Start(":8080"))
}
