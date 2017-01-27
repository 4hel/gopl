package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Movie struct {
	Title  string
	Year   string
	Poster string
	Error  string
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("one argument must be given for title")
	}
	result, err := getMovie(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Poster)
}

func getMovie(title string) (*Movie, error) {
	retval := Movie{}
	movieURL := "http://omdbapi.com/?t=" + url.QueryEscape(os.Args[1]) + "&plot=short&r=json"
	resp, err := http.Get(movieURL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	if err = json.NewDecoder(resp.Body).Decode(&retval); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	if len(retval.Error) > 0 {
		return nil, fmt.Errorf("api returned: %s", retval.Error)
	}
	return &retval, nil
}
