package converter

import (
	"strconv"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// ImageConvert converts the input image file to the specified output format using ffmpeg.
// Parameters:
// - input_path: The path to the input image file.
// - output_path: The path to the output image file.
//
// Returns:
// - nil if the conversion is successful, or an error if the conversion fails.
func ImageConvert(input_path string, output_path string) error {
	err := ffmpeg.Input(input_path).Output(output_path).OverWriteOutput().Run()
	if err != nil {
		return err
	}

	return nil
}

// ImageResize will resize the image to the specified dimensions using ffmpeg
// Parameters:
// - input_path: The path to the input image file.
// - output_path: The path to the output image file.
//
// Returns:
// - nil if the conversion is successful, or an error if the conversion fails.
func ImageResize(input_path string, output_path string, width int, height int) error {
	err := ffmpeg.Input(input_path).
		Filter("scale", ffmpeg.Args{strconv.Itoa(width), strconv.Itoa(height)}).
		Output(output_path).
		OverWriteOutput().
		Run()
	if err != nil {
		return err
	}
	return nil
}

// ImageCompress compresses the image with the specified ratio using ffmpeg
//
// Parameters:
// - input_path: The path to the input image
// - output_path: The path to the output compressed image
// - ratio: The ratio of the image to compress
//
// Returns:
// - nil if the image is compressed successfully, an error otherwise
func ImageCompress(input_path string, output_path string, ratio float64) error {
	err := ffmpeg.Input(input_path).
		Output(output_path, ffmpeg.KwArgs{
			"vf": "scale=iw*" + strconv.FormatFloat(ratio, 'f', -1, 64) + ":ih*" + strconv.FormatFloat(ratio, 'f', -1, 64),
		}).
		OverWriteOutput().
		Run()
	if err != nil {
		return err
	}
	return nil
}
