package main

import (
	"fmt"
	"io"

	"net/http"
)

func main() {
	url := "http://localhost:8080/"

	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error making the request %v\n", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error in reading response body %v\n", err)
		return
	}

	fmt.Printf("Status Code %d\n", resp.StatusCode)
	fmt.Printf("%s", body)

}
