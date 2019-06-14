package images

import (
	"algorithms/dynamic/seamCarving/model"
	"image"
	"image/png"
	"math"
	"os"
)

func Reduce(image image.Image, numberOfSeams int, dimansions model.Dimensions) image.Image {
	for i := 0; i < numberOfSeams; i++ {

		powerArr, tracingArr := GetPowerArray(image, dimansions)
		indexOfMinSeam := findMin(powerArr[len(powerArr)-1])

		image = reduceHorizantally(tracingArr, indexOfMinSeam, image)

	}
	return image

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

		case Left:
			i = i - 1
			j = j - 1
		case Right:
			i = i - 1
			j = j + 1
		case Middle:
			i = i - 1
		case Finish:
			i = i - 1

		}

	}
	return target

}

func CreateImage(img image.Image, path string) {
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
