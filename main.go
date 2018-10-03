package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// type games struct{}

func main() {
	// Request URL
	url := "https://www.giantbomb.com/api/search/?format=json&api_key=0e979a8506def0657887d61aac192b8cefd60eec&query=dragon_age_3&resources=game&field_list=name,deck,guid,image,api_detail_url"

	// Generate Client and define how long it should stay alive.
	brProxyClient := http.Client{
		Timeout: time.Second * 15,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Set headers... I need to better understand headers.
	req.Header.Set("User-Agent", "testing-shit")

	// Make the request to from proxy to API
	res, getErr := brProxyClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	fmt.Println("I'm res.body", res.Body)

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	// JSON returned to server as []byte
	fmt.Println("I'm the pre-unmarshalled body", body)
	// Checking that the JSON is correct.
	fmt.Println("I ought to be JSON man", BytesToString(body))
	games := make(map[string]interface{})
	jsonErr := json.Unmarshal(body, &games)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(games)
	fmt.Println("I'm correctly unmarshalled json fool!", games)

}

func BytesToString(data []byte) string {
	return string(data[:])
}
