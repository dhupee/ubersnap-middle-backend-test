package converter_test

import (
	"image"
	_ "image/jpeg" // Replace with the appropriate image format package
	_ "image/png"
	"os"
	"testing"

	"github.com/dhupee/ubersnap-middle-backend-test/converter"
)

func TestImageConverter(t *testing.T) {
	input_path := "../img/sample_1920x1280.png"
	output_path := "/tmp/ubersnap-test/converted.jpeg"

	// if the output dir does not exist, create it
	if _, err := os.Stat("/tmp/ubersnap-test"); os.IsNotExist(err) {
		os.Mkdir("/tmp/ubersnap-test", 0755)
	}

	// convert the image to jpeg
	err := converter.ImageConvert(input_path, output_path)
	if err != nil {
		t.Error(err)
	}

	// check if the image is a jpeg
	file, err := os.Open(output_path)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	_, format, err := image.DecodeConfig(file)
	if err != nil {
		t.Error(err)
	}
	if format != "jpeg" {
		t.Error("Expected jpeg format, got", format)
	}
}

func TestImageResize(t *testing.T) {
	input_path := "../img/sample_1920x1280.png"
	output_path := "/tmp/ubersnap-test/resized.png"
	width_target := 200
	height_target := 200

	// if the output dir does not exist, create it
	if _, err := os.Stat("/tmp/ubersnap-test"); os.IsNotExist(err) {
		os.Mkdir("/tmp/ubersnap-test/", 0755)
	}

	// resize the image to 200x200
	err := converter.ImageResize(input_path, output_path, width_target, height_target)
	if err != nil {
		t.Error(err)
	}

	// check the width and height of the resized image
	file, err := os.Open(output_path)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	config, _, err := image.DecodeConfig(file)
	if err != nil {
		t.Error(err)
	}

	if config.Width != width_target || config.Height != height_target {
		t.Error("Expected width and height to be", width_target, "and", height_target, "but got", config.Width, "and", config.Height)
	}
}
