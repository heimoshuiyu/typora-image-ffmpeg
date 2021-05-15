package main

import (
	"fmt"
	"os"
	"os/exec"
)

var helpText string = `Typora image ffmpeg compress service
Configure in typora command as typora-image-ffmpeg`

func main() {
	var err error

	// if not input arguments, print help text, and exit program
	args := os.Args[1:]
	if (len(args) == 0) {
		fmt.Println(helpText)
		os.Exit(1)
	}

	for _, v := range args {
		outputFile := v + "_compressed.jpg"

		// run ffmpeg to compress photo
		cmd := exec.Command("ffmpeg", "-i", v, "-y", outputFile)
		err = cmd.Run()
		if err != nil {
			fmt.Println("Failed at compressing", v, err.Error())
			os.Exit(2)
		}

		// delete original file
		os.Remove(v)

		fmt.Println("file://" + outputFile)
	}
}
