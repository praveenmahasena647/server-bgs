package main

import (
	"fmt"
	"os"

	"github.com/praveenmahasena647/bg-server/api"
)

func main() {
	var err error = api.RunServer()

	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(-1)
	}
}
