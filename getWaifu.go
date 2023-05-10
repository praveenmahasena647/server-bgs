package main

import (
	"io/ioutil"
	"net/http"
)

func getHttpWaifuImage() ([]byte, error) {
	var connection, connectionErr = http.Get("https://pic.re/image")
	if connectionErr != nil {
		return nil, connectionErr
	}
	defer connection.Body.Close()

	var data, dataErr = ioutil.ReadAll(connection.Body)

	return data, dataErr
}
