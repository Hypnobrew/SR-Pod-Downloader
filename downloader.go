package main

import (
	"fmt"
    "net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
)

type topnode struct {
	Copyright string
    Podfiles []podfile
	Pagination pagination
}

type podfile struct {
	Url string
	Publishdateutc string
	Title string
}

type pagination struct {
	Page int
	Nextpage string
} 

type podData struct {
	title string
	url string
}

func main() {
	
	metaData := fetchMetaData("2399")
	
	for _, v := range metaData {		
		response, err := download(v.url)
		if err != nil {
			fmt.Println("Could not download data")
			return
		}
		saveToFile(response, v.title)	
	}
}

func fetchMetaData(programID string) ([]podData) {

	url := fmt.Sprintf("http://api.sr.se/api/v2/podfiles?programid=%s&format=json", programID)	
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		fmt.Println("Could not fetch meta data")
		return nil
	}
	
	content, err := ioutil.ReadAll(response.Body)
	var parsedData topnode 
	err = json.Unmarshal(content, &parsedData)
	
	if err != nil {
        fmt.Println("error:", err)
    }
	
	pods := make([]podData, 25)
	
	for i, podfile := range parsedData.Podfiles {
		pods[i] = podData{
			title : strings.Replace(podfile.Title, " ", "_", -1),
			url : podfile.Url,
		}	
        //fmt.Printf("%s %s\n", podfile.Title, podfile.Url)
    }
	return pods
}

func saveToFile(data []byte, title string) {
	err := ioutil.WriteFile(title, data, 0644)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Could not copy data")
		return
	}
	fmt.Println("Saved file")
}

func download(url string) ([]byte, error) {
	response, err := http.Get(url)
	defer response.Body.Close()
    body, err := ioutil.ReadAll(response.Body)
	return body, err
}