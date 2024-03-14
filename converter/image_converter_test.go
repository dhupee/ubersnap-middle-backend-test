package converter_test

import (
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

	// check if the file exists
	if _, err := os.Stat(output_path); os.IsNotExist(err) {
		t.Error(err)
	}
}

// func TestImageResize
