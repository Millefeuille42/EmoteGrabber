package grab

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetFromLink(uri string) (string, error) {
	_, err := url.ParseRequestURI(uri)
	if err != nil {
		return "", err
	}
	u, err := url.Parse(uri)
	if err != nil {
		return "", err
	}
	if !isValidFormat(u) {
		return "", errors.New("image is not in a supported format")
	}
	response, err := http.Get(uri)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}
	return string(body), nil
}
