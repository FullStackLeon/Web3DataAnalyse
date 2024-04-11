package utils

import (
	"io"
	"log"
	"net/http"
)

func GetUrlWithHeader(url string, header map[string]string) (interface{}, error) {
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	for k, v := range header {
		req.Header.Add(k, v)
	}
	res, _ := http.DefaultClient.Do(req)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Wallet query close err", err)
			return
		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Wallet query read err", err)
		return nil, err
	}

	return string(body), nil
}
