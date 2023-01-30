package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func http_request(client *http.Client, body string) {
	req, err := http.NewRequest("POST", "http://127.0.0.1:8080/users", nil)
	if err != nil {
		panic("new request failed")
	}
	req.URL.RawQuery = url.Values{
		"id": {"3"},
	}.Encode()
	// req.Body = ioutil.NopCloser(bytes.NewReader([]byte(body)))
	resp, err := client.Do(req)
	if err != nil {
		panic("client do req failed")
	}
	respData, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("return:%s\n", respData)
}

func main() {
	body := "id=3"
	client := &http.Client{Timeout: time.Duration(20) * time.Second}
	http_request(client, body)

}
