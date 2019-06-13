package images

import (
	"algorithms/dynamic/seamCarving/model"
	"image"
	"image/png"
	"log"
	"math"
	"os"
)

const (
	Finish = -3
	Left   = -1
	Right  = 1
	Middle = 0
)

func init() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
}

func GetPowerArray(image image.Image, dim model.Dimensions) ([][]float64, [][]int) {

	powArr, tracArr := getPowerArray(image, dim)

	return powArr, tracArr

}
func ReadImage(path string) (image.Image, model.Dimensions) {
	imageFile, err := os.Open(path)
	if err != nil {
		log.Printf("Error happend while reading file : %v", err)
		os.Exit(0)
	}
	img, dim, err := getPixelArray(imageFile)
	if err != nil {
		log.Printf("Error happend while getting pixelArray : %v", err)
	}
	return img, dim

}
func getPixelArray(imageFIle *os.File) (image.Image, model.Dimensions, error) {
	imgConfig, _, er := image.DecodeConfig(imageFIle)

	if er != nil {
		log.Printf("ErGetPowerArrayror happend while getting Configs : %v", er)

	}
	dim := model.Dimensions{Height: imgConfig.Height, Width: imgConfig.Width}
	imageFIle.Seek(0, 0)
	img, _, err := image.Decode(imageFIle)
	return img, dim, err
}
func getPowerArray(img image.Image, dimensions model.Dimensions) ([][]float64, [][]int) {
	powerArr := make([][]float64, dimensions.Height)
	tracingArr := make([][]int, dimensions.Height)
	tracingArr[0] = make([]int, dimensions.Width)
	powerArr[0] = make([]float64, dimensions.Width)
	for m := 0; m < dimensions.Width; m++ {

		delta := getPowerAt(0, m, img, dimensions)
		powerArr[0][m] = delta
		tracingArr[0][m] = Finish
	}

	for i := 1; i < dimensions.Height; i++ {
		powerArr[i] = make([]float64, dimensions.Width)
		tracingArr[i] = make([]int, dimensions.Width)

		for j := 0; j < dimensions.Width; j++ {
			delta := getPowerAt(i, j, img, dimensions)
			lPrev := getPrev(powerArr, i-1, j-1, dimensions)
			rPrev := getPrev(powerArr, i-1, j+1, dimensions)
			mPrev := getPrev(powerArr, i-1, j, dimensions)
			power, prev := min(lPrev, mPrev, rPrev)

			powerArr[i][j], tracingArr[i][j] = delta+power, prev

		}
	}

	return powerArr, tracingArr

}
func getPrev(powerArr [][]float64, i, j int, dim model.Dimensions) float64 {
	if i < 0 || j < 0 || j >= dim.Width {
		return math.MaxFloat64
	}
	return powerArr[i][j]
}
func getPowerAt(i, j int, img image.Image, dim model.Dimensions) float64 {
	var r1, g1, b1 uint32
	var r2, g2, b2 uint32
	if j > 0 {
		r1, g1, b1, _ = img.At(j-1, i).RGBA()

	} else {
		r1, b1, g1 = 0, 0, 0
	}
	if j+1 < dim.Width {
		r2, g2, b2, _ = img.At(j+1, i).RGBA()
	} else {
		r2, b2, g2 = 0, 0, 0
	}

	deltaR := float64((r1 - r2) * (r1 - r2))
	deltaG := float64((g1 - g2) * (g1 - g2))
	deltaB := float64((b1 - b2) * (b1 - b2))

	delta := math.Pow(deltaR+deltaG+deltaB, 0.5)
	return delta

}

func min(a, b, c float64) (float64, int) {
	var mini float64
	var prev int
	if a < b {
		mini = a
		prev = Left
	} else {
		mini = b
		prev = Middle
	}
	if mini < c {
		return mini, prev
	}
	return c, Right
}
