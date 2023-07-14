package helpers

import (
	"io"
	"net/http"
)

func HTTPHelper(l string) ([]byte, error) {
	var con, conErr = http.Get(l)

	if conErr != nil {
		return nil, conErr
	}
	defer con.Body.Close()

	var bytes, err = io.ReadAll(con.Body)

	if err != nil {
		return nil, conErr
	}
	return bytes, err
}
