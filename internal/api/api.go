package api

import (
	"bytes"
	"fmt"
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
func (a API) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}

// Post
func (a API) Post(url string, body []byte) string {
	resp, err := http.Post(a.url+url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}
