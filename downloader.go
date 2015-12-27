package main

import (
	"fmt"
	"os"
    "net/http"
	"io/ioutil"
)

func main() {	
	file, err := createFile()
	if err != nil {
		fmt.Println("Could not create file")
		return
	}
	
	response, err := download()
	if err != nil {
		fmt.Println("Could not download data")
		return
	}
	saveToFile(response, file)	
}

func saveToFile(data []byte, file *os.File) {
	err := ioutil.WriteFile("test.mp3", data, 0644)
	if err != nil {
		fmt.Print(err)
		fmt.Println("Could not copy data")
		return
	}
	fmt.Println("Saved file")
}

func download() ([]byte, error) {
	response, err := http.Get("http://sverigesradio.se/topsy/ljudfil/5560699.mp3")
	defer response.Body.Close()
    body, err := ioutil.ReadAll(response.Body)
	return body, err
}

func createFile() (*os.File, error) {
	file, err := os.Create("example.mp3")
	defer file.Close()
	return file, err
}