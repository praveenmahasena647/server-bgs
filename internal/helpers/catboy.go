package helpers

func HTTPGetCatBoy() ([]byte, error) {
	var httpBytes, httpErr = HTTPHelper("https://api.catboys.com/img")

	if httpErr != nil {
		return nil, httpErr
	}
	var link = extractLink(httpBytes)

	var imgBytes, err = HTTPHelper(link)

	if err != nil {
		return nil, err
	}
	return imgBytes, nil
}

func extractLink(b []byte) string {
	var str []byte
	var leng int = len(b)

	for i := 0; i < leng; i++ {
		if b[i] == '\n' {
			continue
		}
		if b[i] == ',' {
			break
		}
		str = append(str, b[i])
	}
	leng = len(str)
	return string(str[8 : leng-1])
}
