package helpers

import (
	"math/rand"
)

func GetImages() ([]byte, error) {

	var (
		randInt  int = rand.Intn(3)
		err      error
		imgBytes []byte = make([]byte, 1)
	)

	switch randInt {
	case 0:
		imgBytes, err = HTTPGetWaifu() // This is done
	case 1:
		imgBytes, err = HTTPGetNeko()
	case 2:
		imgBytes, err = HTTPGetCatBoy()
	}

	return imgBytes, err
}
