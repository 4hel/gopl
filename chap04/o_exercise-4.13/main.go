package main

import (
	"fmt"
	"net/url"
	"os"
)

type Movie struct {
}

func main() {
	movieURL := "http://omdbapi.com/?t=" + url.QueryEscape(os.Args[1]) + "&plot=short&r=json"
	fmt.Println(movieURL)
}
