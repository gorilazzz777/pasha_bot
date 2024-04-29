package img

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
	gim "pasha_bot/pkg/img/imagemerge"
)

func MergeImages(photoPath1, photoPath2, photoResultPath string) {
	grids := []*gim.Grid{
		{ImageFilePath: photoPath1},
		{ImageFilePath: photoPath2},
	}

	// merge the images into a 2x1 grid
	rgba, err := gim.New(grids, 2, 1).Merge()
	if err != nil {
		fmt.Println(err)
	}

	// save the output to jpg or png
	file, err := os.Create(photoResultPath)
	err = jpeg.Encode(file, rgba, &jpeg.Options{Quality: 80})
	if err != nil {
		fmt.Println(err)
	}
}

func DeleteImage(path string) {
	e := os.Remove(path)
	if e != nil {
		log.Fatal(e)
	}
}
