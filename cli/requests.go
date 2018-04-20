package main

import (
	"io"
	"net/http"
)

func addRequestHeaders(req *http.Request, user, secret string) {
	signature := ""
	req.Header.Add("X-C10-Date", "")
	req.Header.Add("X-C10-User", user)
	req.Header.Add("X-C10-Signature", signature)
}

func doPostRequest(url string, in io.Reader, user, secret string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, nil)

	addRequestHeaders(req, user, secret)

	client := &http.Client{}
	resp, err := client.Do(req)

	return resp, err
}
