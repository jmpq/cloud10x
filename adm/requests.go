package main

import (
	"fmt"
	"github.com/jmpq/cloud10x/v1"
	//"io"
	"net/http"
	"strings"
	"time"
)

func addRequestHeaders(req *http.Request, data string, user, secret string) {
	iso8601Format := "20060102T150405Z"
	now := time.Now().UTC()
	date := now.Format(iso8601Format)

	signature := v1.Signature(date, data, secret)

	req.Header.Add("X-C10X-Date", date)
	req.Header.Add("X-C10X-User", user)
	req.Header.Add("X-C10X-Signature", signature)
	fmt.Printf("Signature %s\n", signature)
}

func doPostRequest(url string, data string, user, secret string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(data))

	addRequestHeaders(req, data, user, secret)

	client := &http.Client{}
	resp, err := client.Do(req)

	return resp, err
}

func doGetRequest(url string, user, secret string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)

	addRequestHeaders(req, "", user, secret)

	client := &http.Client{}
	resp, err := client.Do(req)

	return resp, err
}
