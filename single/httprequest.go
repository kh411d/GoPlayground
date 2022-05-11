package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	Client := &http.Client{
		Timeout: 30 * time.Second,
	}

	url := "http://0.0.0.0:3007/v1/health_check"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := Client.Do(req)
	if err != nil {
		fmt.Printf("Error Client Do: %v Res : %v\n", err, res)
	}

	if res != nil {
		defer res.Body.Close()

		fmt.Printf("Response: %v\n", res)

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("Error ReadAll: %v\n", err)
		}

		r := map[string]interface{}{}
		err = json.Unmarshal(body, &r)
		if err != nil {
			fmt.Printf("Error Unmarshal: %v\n", err)
		}

		fmt.Printf("Response Unmarshal: %v\n", r)
	}

}
