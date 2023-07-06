package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func fetchContent(url string, contentChannel chan<- map[string]string, errChannel chan<- error) {
	resp, err := http.Get(url)
	if err != nil {
		errChannel <- fmt.Errorf("failed to fetch URL %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errChannel <- fmt.Errorf("failed to read response body for URL %s: %v", url, err)
		return
	}

	content := map[string]string{
		url: string(body),
	}

	contentChannel <- content
}

func Task_5_solution() {
	urls := []string{
		"https://api.chucknorris.io/jokes/random",
		"https://api.chucknorris.io/jokes/search?query=9",
		"https://api.chucknorris.io/jokes/random?category=test",
		"https://api.chucknorris.io/jokes/random?category=dev",
	}

	contentChannel := make(chan map[string]string)
	errChannel := make(chan error)

	for _, url := range urls {
		go fetchContent(url, contentChannel, errChannel)
	}

	contentMap := make(map[string]string)

	for range urls {
		select {
		case content := <-contentChannel:
			for key, value := range content {
				contentMap[key] = value
			}
		case err := <-errChannel:
			fmt.Printf("Error: %v\n", err)
		}
	}

	// Print the fetched content for each URL
	for url, content := range contentMap {
		fmt.Printf("URL: %s\nContent:\n%s\n\n", url, content)
	}
}
