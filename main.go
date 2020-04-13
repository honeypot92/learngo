package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

// func handleScrape(c echo.Context) error {
// defer os.Remove(fileName)
// term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
// scrapper.Scrape(term)
// return c.Attachment(fileName, fileName)
// }

func main() {
	e := echo.New()
	// g := e.Group("/admin")

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${method} ${error} ${path} ${latency}" + "\n",
	}))
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handleHome)
	// e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":8000"))
	//	scrapper.Scrape("term")
}
