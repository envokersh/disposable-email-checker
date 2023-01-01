package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
)

const DEFAULT_LIST = "https://disposable.github.io/disposable-email-domains/domains.txt"

func main() {
	list_url, ok := os.LookupEnv("DOMAIN_LIST")
	if !ok {
		list_url = DEFAULT_LIST
	}

	s := gocron.NewScheduler(time.UTC)
	var domains []string

	s.Every(24).Hours().Do(func() {
		resp, err := http.Get(list_url)
		if err != nil {
			log.Fatalln(err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		domains = strings.Split(string(body), "\n")
	})

	go s.StartAsync()

	app := fiber.New()
	app.All("/check/:email", func(c *fiber.Ctx) error {
		email := c.Params("email")

		parts := strings.Split(email, "@")
		if len(parts) < 2 {
			return c.JSON(fiber.Map{
				"success": false,
				"message": "Invalid email",
			})
		}

		res := slices.Contains(domains, parts[1])

		return c.JSON(fiber.Map{
			"success":    true,
			"disposable": res,
		})
	})

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
