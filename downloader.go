package main

import (
	"fmt"
    "net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
	"flag"
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
	
	programID := flag.Int("id", 0, "The id of the program you want to download podcasts for")
	numberOfPods := flag.Int("pods", 10, "The number of podcasts to download")	
	flag.Parse()
	
	if(*programID == 0) {
		fmt.Println("Error: Please specify the id on the program you wish to download. Use --help for more information.")
		return
	}
	
	metaData := prepareMetaData(*programID, *numberOfPods)
	for i, v := range metaData {
		fmt.Printf("%d: Downloading file %s", i, v.url)
		response, err := download(v.url)
		if err != nil {
			fmt.Printf("Error downloading %s",  v.url)
			return
		}
		saveToFile(response, v.title)	
	}
}

func prepareMetaData(programID int, numberOfPods int) ([]podData) {
	
	var allPodcasts []podData
	url := fmt.Sprintf("http://api.sr.se/api/v2/podfiles?programid=%d&format=json", programID)
	for haveMorePages := true; haveMorePages; haveMorePages = (url != "") {
		pagedPodcasts, urlToNextPage := fetchMetaData(url)
		allPodcasts = append(allPodcasts, pagedPodcasts...)
		
		if(len(allPodcasts) >= numberOfPods) {
			limitedPodcastas := make([]podData, numberOfPods)
			copy(limitedPodcastas, allPodcasts)
			allPodcasts = limitedPodcastas
			break
		}
		url = urlToNextPage	
	}
	return allPodcasts
}

func fetchMetaData(url string) ([]podData, string) {
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		fmt.Println("Could not fetch meta data")
		return nil, ""
	}
	
	content, err := ioutil.ReadAll(response.Body)
	var parsedData topnode 
	err = json.Unmarshal(content, &parsedData)
	
	if err != nil {
        fmt.Println("error:", err)
    }
	
	pods := make([]podData, len(parsedData.Podfiles))
	for i, podfile := range parsedData.Podfiles {
		pods[i] = podData{
			title : strings.Replace(podfile.Title, " ", "_", -1),
			url : podfile.Url,
		}	
    }
	return pods, parsedData.Pagination.Nextpage
}

func saveToFile(data []byte, title string) {
	err := ioutil.WriteFile(title, data, 0644)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Could not copy data")
		return
	}
}

func download(url string) ([]byte, error) {
	response, err := http.Get(url)
	defer response.Body.Close()
    body, err := ioutil.ReadAll(response.Body)
	return body, err
}