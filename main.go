package main

import (
	"errors"
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request Failed")

func main() {
	results := make(map[string]string)
	c := make(chan result)
	urls := []string{
		"https://www.facebook.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://www.instagram.com/",
		"https://soundcloud.com/",
		"https://www.naver.com/",
		"https://www.zum.com/",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		rslt := <-c
		results[rslt.url] = rslt.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- result) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- result{url: url, status: status}
}
