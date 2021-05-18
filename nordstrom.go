package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/nickname32/discordhook"
)

func main() {
	fmt.Println("here")
	wa, err := discordhook.NewWebhookAPI(sample, "smple-FY", true, nil)
	if err != nil {
		panic(err)
	}

	wh, err := wa.Get(nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(wh.Name)

	msg, err := wa.Execute(nil, &discordhook.WebhookExecuteParams{
		Content: "Example text",
		Embeds: []*discordhook.Embed{
			{
				Title:       "Hi there",
				Description: "This is description",
			},
		},
	}, nil, "")
	if err != nil {
		panic(err)
	}

	fmt.Println(msg.ID)

	wh, err = wa.Modify(nil, &discordhook.WebhookModifyParams{
		Name: "This is a new default webhook name",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(wh)

	err = wa.Delete(nil)
	if err != nil {
		panic(err)
	}
}

//webhook https://discord.com/api/webhooks//
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
