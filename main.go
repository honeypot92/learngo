package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	var results = make(map[string]string)
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
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "failed"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking Url:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}
