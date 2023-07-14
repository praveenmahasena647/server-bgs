package helpers

func HTTPGetNeko() ([]byte, error) {
	var httpBytes, httpErr = HTTPHelper("https://nekos.best/api/v2/neko")

	if httpErr != nil {
		return nil, httpErr
	}
	var link = extractNekoLink(httpBytes)

	var imgBytes, err = HTTPHelper(link)

	if err != nil {
		return nil, err
	}
	return imgBytes, nil

}

func extractNekoLink(b []byte) string {
	var url string
	var leng int = len(b) - 5

	for leng >= 0 {
		if b[leng] == '"' {
			break
		} else {
			url = string(b[leng]) + url
		}
		leng--
	}
	return url
}
