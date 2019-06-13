package main

import (
	"algorithms/dynamic/seamCarving/images"
	"algorithms/dynamic/seamCarving/model"
	"fmt"
	"image"
	"image/png"
	"math"
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
	image = reduce(image, horizantalSeams, dimensions)
	/***
	we flip the image and then reduce again in the horizental direction
	we get the vertical reduction
	***/
	newDimensions := model.Dimensions{Height: dimensions.Width, Width: dimensions.Height}
	image = reduce(imaging.Rotate90(image), verticalSeams, newDimensions)

	/***
	we need to flip the image to its initial position before saving the result

	***/
	createImage(imaging.Rotate270(image), outputFile)

}
func reduce(image image.Image, numberOfSeams int, dimansions model.Dimensions) image.Image {
	for i := 0; i < numberOfSeams; i++ {

		powerArr, tracingArr := images.GetPowerArray(image, dimansions)
		indexOfMinSeam := findMin(powerArr[len(powerArr)-1])

		image = reduceHorizantally(tracingArr, indexOfMinSeam, image)

	}
	return image

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

func reduceHorizantally(tracingArr [][]int, index int, img image.Image) image.Image {

	height, width := len(tracingArr), len(tracingArr[0])
	bounds := img.Bounds()
	target := image.NewNRGBA(image.Rect(0, 0, bounds.Dx()-1, bounds.Dy()))
	i, j := height-1, index

	for i >= 0 {
		k := 0
		for m := 0; m < width; m++ {

			if m != j {
				target.Set(k, i, img.At(m, i))
				k++
			}
		}

		switch tracingArr[i][j] {

		case images.Left:
			i = i - 1
			j = j - 1
		case images.Right:
			i = i - 1
			j = j + 1
		case images.Middle:
			i = i - 1
		case images.Finish:
			i = i - 1

		}

	}
	return target

}

func createImage(img image.Image, path string) {
	fg, err := os.Create(path)
	defer fg.Close()
	if err != nil {
		panic(err)
	}
	err = png.Encode(fg, img)
	if err != nil {
		panic(err)
	}

}
func findMin(arr []float64) int {
	min := math.MaxFloat64
	index := -1

	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
			index = i
		}
	}

	return index
}
