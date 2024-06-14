package main

import (
	"log"

	"github.com/HammiAhlan/ws-hammi2024/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	_ "github.com/gofiber/swagger/example/docs"
	"github.com/HammiAhlan/ws-hammi2024/url"

	"github.com/gofiber/fiber/v2"
	
)

// @title TES SWAGGER ULBI
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/HammiAhlan
// @contact.email 714220062@std.ulbi.ac.id

// @host ws-allan2024-0d01e8eb9e77.herokuapp.com
// @BasePath /
// @schemes https http

func main() {
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
