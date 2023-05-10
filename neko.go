package main

import (
	"io/ioutil"
	"net/http"
)

func getHttpNekoImage() ([]byte, error) {
	var err error
	urlByte, err := nekoUrl()

	if err != nil {
		return nil, err
	}
	return httpGetNekoImageBytes(urlByte)
}

func httpGetNekoImageBytes(urlByte []byte) ([]byte, error) {
	var url string = extractLinkNeko(urlByte)

	var conn, connErr = http.Get(url)

	if connErr != nil {
		return nil, connErr
	}
	defer conn.Body.Close()

	var imgBytes, byteErr = ioutil.ReadAll(conn.Body)

	return imgBytes, byteErr
}

func extractLinkNeko(urlByte []byte) string {
	var url string
	var leng int = len(urlByte) - 5

	for leng >= 0 {
		if urlByte[leng] == '"' {
			break
		} else {
			url = string(urlByte[leng]) + url
		}
		leng--
	}

	return url
}

func nekoUrl() ([]byte, error) {
	// done
	var conn, connErr = http.Get("https://nekos.best/api/v2/neko")

	if connErr != nil {
		return nil, connErr
	}

	defer conn.Body.Close()

	urlByte, err := ioutil.ReadAll(conn.Body)

	return (urlByte), err
}
