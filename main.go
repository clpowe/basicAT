package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
    return c.Render("index", fiber.Map{
        "hello": "world",
    },"layouts/main");
})

app.Post("/api/outliers", func(c *fiber.Ctx) error {
		max, avg , sd := outliers(c)

		return c.Render("partials/outliers", fiber.Map{
			"Avg": fmt.Sprintf("%.2f", avg) ,
			"SD":  fmt.Sprintf("%.2f", sd) ,
			"Max": fmt.Sprintf("%.2f", max) ,
		})
	})

	log.Fatal(app.Listen(":8080"))
}

func outliers(c *fiber.Ctx) (float64,float64,float64,)  {
		cavg, _ := strconv.ParseFloat(c.FormValue("cavg"), 64)
		cSD, _ := strconv.ParseFloat(c.FormValue("cSD"), 64)
		vavg, _ := strconv.ParseFloat(c.FormValue("vavg"), 64)
		vSD, _ := strconv.ParseFloat(c.FormValue("vSD"), 64)

		avg := (cavg + vavg) / 2
		sd := (cSD + vSD) / 2

		max := (sd * 3) + avg

		return max, avg, sd	
}