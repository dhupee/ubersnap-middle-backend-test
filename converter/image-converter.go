package process

import (
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
