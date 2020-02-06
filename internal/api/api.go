package api

import (
	"bytes"
	"github.com/pniedzwiedzinski/photoprism-cli/internal/utils"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// API - communicate with photoprism api
type API struct {
	url   string
	token string
}

// NewAPI constructor
func NewAPI(ip string) API {
	url := "http://" + ip + ":2342/api/v1/"
	tokenFile := utils.GetTokenLocation() + "token.txt"
	data, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		return API{url: url, token: ""}
	}
	return API{url: url, token: string(data)}
}

// Create http.Request, add token, and send to the url with body
func (a API) doRequest(method string, url string, body string) (*http.Response, error) {
	client := http.Client{}
	req, err := http.NewRequest(method, a.url+url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Session-Token", a.token)
	if method == "POST" {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetToken - return current session token
func (a API) GetToken() string {
	return a.token
}

// Get - GET HTTP request to `http://server:2342/api/v1/{url}`
func (a API) Get(url string) (string, error) {
	resp, err := a.doRequest("GET", url, "")
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

// Post - POST HTTP request to `http://server:2342/api/v1/{url}` with body
func (a API) Post(url string, body string) (string, error) {
	resp, err := a.doRequest("POST", url, body)
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
