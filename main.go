package main

import (
	"log"
	"os"

	process "github.com/dhupee/ubersnap-middle-backend-test/converter"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// fmt.Println("Hello, World!")

	// Make temporary directory if it doesn't exist
	os.Mkdir("/tmp/ubersnap-backend", 0775)

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env")
		log.Println("Using default environment variables")
	}

	// Get the port from the environment variables
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("PORT environment variable is not set")
	}

	app := fiber.New()
	v1 := app.Group("/v1")

	v1.Get("/", RootHandler)
	v1.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// Route to receive file
	v1.Post("/upload", func(c *fiber.Ctx) error {
		// Parse the form file
		image, err := c.FormFile("image")
		if err != nil {
			return err
		}

		tmpDir := "/tmp/ubersnap-backend/"

		// Save the file to the temporary directory
		os.Create(tmpDir + image.Filename)

		// Run a function to process the file
		err = process.ImageConvert(image.Filename, "test.png")
		if err != nil {
			return err
		}

		return c.SendString("File uploaded and processed successfully")
	})

	log.Fatal(app.Listen(":" + PORT))
}

func RootHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
	// return c.SendFile("./assets/welcome.txt")
}
