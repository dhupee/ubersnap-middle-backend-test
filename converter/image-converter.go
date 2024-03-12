package process

import (
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// ImageConvert converts the input image file to the specified output format using ffmpeg.
// It takes the input file path and the output file path as arguments.
// If the conversion is successful, it returns nil. Otherwise, it returns an error.
func ImageConvert(input_path string, output_path string) error {
	err := ffmpeg.Input(input_path).Output(output_path).OverWriteOutput().Run()
	if err != nil {
		return err
	}

	return nil
}
