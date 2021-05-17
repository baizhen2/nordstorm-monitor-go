package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	json1 := GetRequest("https://www.nordstromrack.com/api/search2/catalog/search?query=robot")

	fmt.Println(GetNumItems(json1))
}

type nordstrom_data struct {
	Search_cluster struct {
		Num_found int `json:"num_found"`
	}
}

func GetRequest(requestURL string) string {
	resp, err := http.Get(requestURL)

	if err != nil {
		return "Error has occured"
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return string(body)
}

func GetNumItems(body string) int {
	textBytes := []byte(body)

	data := nordstrom_data{}

	err := json.Unmarshal(textBytes, &data)

	if err != nil {
		fmt.Println(err)
		return 999999999
	}

	return data.Search_cluster.Num_found
}
