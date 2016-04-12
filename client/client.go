package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Post(url string, data string, username string, password string) (string, string) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))

	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)

	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.Status, string(body))
	return resp.Status, string(body)
}
