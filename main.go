package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	converter "github.com/dhupee/ubersnap-middle-backend-test/converter"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// fmt.Println("Hello, World!")

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

	app := fiber.New(fiber.Config{ // config for the server
		BodyLimit: 10 * 1024 * 1024, // 10 MB
	})

	v1 := app.Group("/v1")

	v1.Get("/", RootHandler)
	v1.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
	v1.Post("/convert", ConvertHandler)

	// Start Fiber server
	log.Fatal(app.Listen(":" + PORT))
}

func RootHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
	// return c.SendFile("./assets/welcome.txt")
}

// Route to receive file
func ConvertHandler(c *fiber.Ctx) error {
	// Headers
	fileTarget := c.Get("file-target")

	// Make temporary directory if it doesn't exist
	tmpDir := "/tmp/ubersnap-backend"
	if _, err := os.Stat(tmpDir); os.IsNotExist(err) {
		os.Mkdir(tmpDir, 0775)
	}

	// This will be passed in the body of the request
	image, err := c.FormFile("image")
	if err != nil {
		log.Println(err)
		return err
	}

	// TODO: if the file is not an image, return an error
	// TODO: if the fileTarget is not in the list, also return an error

	// extract image name without extension
	imageName := strings.Split(image.Filename, ".")[0]
	imageType := strings.Split(image.Filename, ".")[1]
	log.Println(imageName) // comment this once you dont need these
	log.Println(imageType)
	log.Println(fileTarget)

	// if target directory doesn't exist, create it
	if _, err := os.Stat(fmt.Sprintf(tmpDir + "/" + imageName)); os.IsNotExist(err) {
		os.Mkdir(tmpDir+"/"+imageName, 0775)
	}
	// TODO:add else if if there's similar directory, add "directory-1" or "directory-n"

	// Save file to root directory:
	err = c.SaveFile(image, fmt.Sprintf(tmpDir+"/"+imageName+"/input."+imageType))
	if err != nil {
		log.Println(err)
		return err
	}

	inputPath := tmpDir + "/" + imageName + "/input." + imageType
	outputPath := tmpDir + "/" + imageName + "/output." + fileTarget

	err = converter.ImageConvert(inputPath, outputPath)
	if err != nil {
		log.Println(err)
		return err
	}

	// TODO: add a route to download the file
	return c.SendString("File uploaded and saved successfully")
}
