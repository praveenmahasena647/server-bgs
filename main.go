package main

func main() {
	var err error
	err = runServer()

	if err != nil {
		panic(err)
	}
}
