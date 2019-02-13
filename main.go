package main
import (
    // "io"
    // "net/http"
    // "os"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

func main()  {
	e := echo.New()

    e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
    e.File("/", "public/index.html")
    e.Static("/assets", "public/assets")
    e.Logger.Fatal(e.Start(":9000"))
}