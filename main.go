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

	// Make temporary directory if it doesn't exist
	tmpDir := "/tmp/ubersnap-backend"
	os.Mkdir(tmpDir, 0775)

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

	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024, // 10 MB
	})

	v1 := app.Group("/v1")

	v1.Get("/", RootHandler)
	v1.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// Route to receive file
	v1.Post("/convert", func(c *fiber.Ctx) error {
		fileTarget := c.Get("file-target")

		// Parse the form file
		image, err := c.FormFile("image")
		if err != nil {
			log.Println(err)
			return err
		}

		// extract image name without extension
		imageName := strings.Split(image.Filename, ".")[0]
		imageType := strings.Split(image.Filename, ".")[1]
		log.Println(imageName)
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

		return c.SendString("File uploaded and saved successfully")
		// filename := image.Filename
		//
		// tmpDir := ("/tmp/ubersnap-backend/" + filename + "/")
		//
		// // Save the file to the temporary directory
		// // os.Create(tmpDir + "original")
		// c.SaveFile(image, tmpDir+"original.png")
		//
		// // // Run a function to process the file
		// // err = process.ImageConvert(image.Filename, "test.png")
		// // if err != nil {
		// // 	return err
		// // }
		//
		// return c.SendString("File uploaded and processed successfully")
	})

	// Start Fiber server
	log.Fatal(app.Listen(":" + PORT))
}

func RootHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
	// return c.SendFile("./assets/welcome.txt")
}
