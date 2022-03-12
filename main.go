package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func handleHome(c echo.Context) error {
	return c.File("Home.html")
}

func handleScrape(c echo.Context) error {
	fmt.Println(c.FormValue("term"))

	return nil
}

func main() {
	// scrapper.Scrape("blockchain")

	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
