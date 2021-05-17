package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	Monitor("robot")
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

func Monitor(request string) {
	monitor := true

	default_endpoint := "https://www.nordstromrack.com/api/search2/catalog/search?query="

	query := strings.Replace(request, " ", "+", -1)

	full_query := default_endpoint + query

	for {

		json1 := GetRequest(full_query)

		fmt.Println(GetNumItems(json1))

		time.Sleep(5000 * time.Millisecond)

		if monitor == false {
			break
		}
	}
}
