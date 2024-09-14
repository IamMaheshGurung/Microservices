package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	formUrl := "http://localhost:8080/form"

	var name string
	var email string
	fmt.Print("Enter your name: ")
	_, err := fmt.Scanf("%s\n", &name)
	if err != nil {
		fmt.Println("Error reading name:", err)
		return
	}

	fmt.Print("Enter your email: ")
	_, err = fmt.Scanf("%s\n", &email)
	if err != nil {
		fmt.Println("Error reading email:", err)
		return
	}

	formData := url.Values{
		"name":  {name},
		"email": {email},
	}

	resp, err := http.PostForm(formUrl, formData)

	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Printf("%s", body)

}
