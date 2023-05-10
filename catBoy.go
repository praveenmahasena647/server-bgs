package main

import (
	"io/ioutil"
	"net/http"
)

func getHttpCatBoyImage() ([]byte, error) {
	var con, conErr = http.Get("https://api.catboys.com/img")

	if conErr != nil {
		return nil, conErr
	}
	defer con.Body.Close()

	var linkByte, linkErr = ioutil.ReadAll(con.Body)

	if linkErr != nil {
		return nil, linkErr
	}
	var urlStr = httpCatBoyLink(linkByte)

	return catBoyImageBytes(urlStr)
}

func catBoyImageBytes(urlStr string) ([]byte, error) {
	var con, connErr = http.Get(urlStr)
	if connErr != nil {
		return nil, connErr
	}

	defer con.Body.Close()

	return ioutil.ReadAll(con.Body)
}

func httpCatBoyLink(linkByte []byte) string {
	var str []byte
	var leng int = len(linkByte)

	for i := 0; i < leng; i++ {
		if linkByte[i] == '\n' {
			continue
		}
		if linkByte[i] == ',' {
			break
		}
		str = append(str, linkByte[i])
	}
	leng = len(str)
	return string(str[8 : leng-1])
}
