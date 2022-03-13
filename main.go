package main

import (
	"os"

	"github.com/hyun2/jobScrapper/scrapper"
	"github.com/labstack/echo/v4"
)

const fileName = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("Home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	// fmt.Println(c.FormValue("term"))
	term := c.FormValue("term")
	scrapper.Scrape(term)

	return c.Attachment(fileName, fileName)
}

func main() {
	// scrapper.Scrape("blockchain")

	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
