package helpers

func HTTPGetWaifu() ([]byte, error) {
	var bytes, err = HTTPHelper("https://pic.re/image")
	return bytes, err
}
