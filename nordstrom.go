package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/andersfylling/snowflake"
	"github.com/nickname32/discordhook"
)

func main() {
	fmt.Println("Enter search query: ")
	var query string
	fmt.Scanln(&query)

	fmt.Println("Enter webhookURL: ")
	var hook string
	fmt.Scanln(&hook)

	Monitor(query, hook)
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

func Monitor(request string, hook string) {
	monitor := true

	default_endpoint := "https://www.nordstromrack.com/api/search2/catalog/search?query="

	query := strings.Replace(request, " ", "+", -1)

	full_query := default_endpoint + query

	num_items := 0

	for {

		json1 := GetRequest(full_query)

		if GetNumItems(json1) != num_items {
			num_items = GetNumItems(json1)

			execute(num_items, hook)
		}

		time.Sleep(5000 * time.Millisecond)
		fmt.Println("Working")

		if monitor == false {
			break
		}
	}
}

func CreateWebhook(webhookURL string) (*discordhook.WebhookAPI, error) {
	split := strings.Split(webhookURL, "/")
	flake := snowflake.ParseSnowflakeString(split[5])
	token := split[6]

	return discordhook.NewWebhookAPI(flake, token, true, nil)
}

func execute(numItems int, webhookURL string) {
	hook, err := CreateWebhook(webhookURL)
	if err != nil {
		panic(err)
	}

	wh, err := hook.Get(nil)
	if err != nil {
		panic(err)
	}

	msg, err := hook.Execute(nil, &discordhook.WebhookExecuteParams{
		Embeds: []*discordhook.Embed{
			{
				Title:       "Items on page",
				Description: strconv.Itoa(numItems),
			},
		},
	}, nil, "")
	if err != nil {
		panic(err)
	}

	fmt.Println(msg.ID)
	fmt.Println(wh)
}
