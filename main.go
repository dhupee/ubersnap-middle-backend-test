package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	converter "github.com/dhupee/ubersnap-middle-backend-test/converter"
	utils "github.com/dhupee/ubersnap-middle-backend-test/utils"

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

	// Routes
	v1.Get("/", RootHandler)
	v1.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
	v1.Post("/convert", ConvertHandler)

	// // NOTE: This one is working but can be improved for the serving purposes
	// v1.Get("/download/", func(c *fiber.Ctx) error {
	// 	imageName := c.Get("image-name")
	// 	fileTarget := c.Get("file-target")
	// 	if imageName == "" || fileTarget == "" {
	// 		return c.Status(400).SendString("Missing image name or file target")
	// 	}
	// 	tmpDir := "/tmp/ubersnap-backend"
	// 	log.Println("download: ", tmpDir+"/"+imageName+"/output."+fileTarget)
	// 	return c.SendFile(tmpDir + "/" + imageName + "/output." + fileTarget)
	// })

	// Start Fiber server
	log.Fatal(app.Listen(":" + PORT))
}

func RootHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
	// return c.SendFile("./assets/welcome.txt")
}

// Route to receive file
func ConvertHandler(c *fiber.Ctx) error {
	// TODO: Move this function somewhere

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

	// Accepted file types
	// TODO: check ffmpeg again for accepted file types
	imageTypeList := []string{"jpg", "jpeg", "png", "webp"}

	// extract image name without extension
	imageName := strings.Split(image.Filename, ".")[0]
	imageType := strings.Split(image.Filename, ".")[1]
	log.Println(imageName) // comment this once you dont need these
	log.Println(imageType)
	log.Println(fileTarget)

	// Throw an error if the imageType and fileTarget is not in the list
	// NOTE: In the future we can add API to scan for malware or something
	if !utils.IsInSlice(imageType, imageTypeList) {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid file type", "data": nil})
	}
	if !utils.IsInSlice(fileTarget, imageTypeList) {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid file target", "data": nil})
	}

	// if target directory doesn't exist, create it
	// TODO: Have better naming convention
	if _, err := os.Stat(fmt.Sprintf(tmpDir + "/" + imageName)); os.IsNotExist(err) {
		os.Mkdir(tmpDir+"/"+imageName, 0775)
	}

	// Save the file to the temporary directory
	err = c.SaveFile(image, fmt.Sprintf(tmpDir+"/"+imageName+"/input."+imageType))
	if err != nil {
		return err
	}

	inputPath := tmpDir + "/" + imageName + "/input." + imageType
	outputPath := tmpDir + "/" + imageName + "/output." + fileTarget

	err = converter.ImageConvert(inputPath, outputPath)
	if err != nil {
		log.Println(err)
		return err
	}

	// curlDownloadLink := fmt.Sprintf(`curl --request GET --url http://127.0.0.1:8080/v1/download/ --header 'file-target: %s' --header 'image-name: %s'`, fileTarget, imageName)
	// return c.SendString("File successfully converted, please download by curl: " + curlDownloadLink)
	// return c.SendString("File uploaded and saved successfully")

	// Send the output file
	return c.SendFile(tmpDir + "/" + imageName + "/output." + fileTarget)

	// NOTE: this is not ideal for production code,
	// the correct way is to send the output file to storage bucket like S3 or GCS
}
