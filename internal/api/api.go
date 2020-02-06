package api

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// API bla ble Comment
type API struct {
	url   string
	token string
}

// NewAPI constructor
func NewAPI(ip string) API {
	url := "http://" + ip + ":2342/api/v1/"
	token := ""
	return API{url: url, token: token}
}

// Get
func (a API) Get(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Post
func (a API) Post(url string, body string) (string, error) {
	resp, err := http.Post(a.url+url, "application/json", bytes.NewBuffer([]byte(body)))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
