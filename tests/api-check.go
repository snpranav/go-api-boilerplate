package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	response, _ := http.Get("http://localhost:3000/")
	//     expected_response := "Hello World! The Go API is running inside a contai
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	str_body := string(body)
	if str_body == "Hello World! The Go API is running inside a container using the go-api-boilerplate docker image.\nPlease visit - https://hub.docker.com/r/snpranav/go-api-boilerplate." {
		fmt.Println("Test Passed!")
	} else {
		fmt.Println("Test Failed! Exiting.")
		os.Exit(127)
	}
}
