package main

import (
	"math/rand"
)

func getImages() ([]byte, error) {
	var v = rand.Intn(3)
	var err error
	var imgdata = []byte{}

	if v == 0 {
		imgdata, err = getHttpWaifuImage()
	} else if v == 1 {
		imgdata, err = getHttpNekoImage()
	} else {
		imgdata, err = getHttpCatBoyImage()
	}

	return imgdata, err
}
