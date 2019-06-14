package main

import (
	"algorithms/dynamic/seamCarving/images"
	"algorithms/dynamic/seamCarving/model"
	"fmt"
	"os"
	"strconv"

	"github.com/disintegration/imaging"
)

func main() {
	CommandlineArgs := os.Args[1:]
	if len(CommandlineArgs) < 4 {
		fmt.Println("Missing argumnets input,output,width,height!")
		os.Exit(0)
	}
	inputFile, outputFile, newWidth, newHeight := getCommandLinesArgs(CommandlineArgs)
	image, dimensions := images.ReadImage(inputFile)

	horizantalSeams := dimensions.Width - newWidth
	verticalSeams := dimensions.Height - newHeight
	/***
	reduces in the horizontal direction
	***/
	image = images.Reduce(image, horizantalSeams, dimensions)
	/***
	we flip the image and then reduce again in the horizental direction
	we get the vertical reduction
	***/
	newDimensions := model.Dimensions{Height: dimensions.Width, Width: dimensions.Height}
	image = images.Reduce(imaging.Rotate90(image), verticalSeams, newDimensions)

	/***
	we need to flip the image to its initial position before saving the result

	***/
	images.CreateImage(imaging.Rotate270(image), outputFile)

}

func getCommandLinesArgs(CommandlineArgs []string) (string, string, int, int) {
	inputFile := CommandlineArgs[0]
	outputFile := CommandlineArgs[1]
	newWidth, err := strconv.Atoi(CommandlineArgs[2])
	if err != nil {
		fmt.Printf("error with converting new width parameter :%v", err)
		os.Exit(0)

	}
	newHeight, err := strconv.Atoi(CommandlineArgs[3])
	if err != nil {
		fmt.Printf("error with converting new Height parameter :%v", err)
		os.Exit(0)
	}
	return inputFile, outputFile, newWidth, newHeight

}
