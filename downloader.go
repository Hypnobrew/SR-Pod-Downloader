package main

import (
	"fmt"
	"os"
    "net/http"
    "io"
)

func main() {	
	file, err := createFile()
	if err != nil {
		fmt.Println("Could not create file")
	}
	
	response, err := download()
	if err != nil {
		fmt.Println("Could not download data")
	}
	
	saveToFile(response, file)	
}

func saveToFile(response *http.Response, file *os.File) {
	size, err := io.Copy(file, response.Body)
	if err == nil {
		fmt.Println("Could not create file")
		return
	}
	
	fmt.Printf("Did write %d bytes to file", size)
}

func download() (*http.Response, error) {
	response, err := http.Get("http://example.com/")
	defer response.Body.Close()
	return response, err
}

func createFile() (*os.File, error) {
	file, err := os.Create("example.mp3")
	defer file.Close()
	return file, err
}