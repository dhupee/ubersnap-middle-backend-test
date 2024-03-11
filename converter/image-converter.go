package process

import (
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func ImageConvert(input_path string, output_path string) error {
	err := ffmpeg.Input(input_path).Output(output_path).OverWriteOutput().Run()
	if err != nil {
		return err
	}

	return nil
}
