package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchUrlData(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error getting data from url: %s err: %w", url, err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error getting response body, err: %w", err)
	}

	return body, nil
}
