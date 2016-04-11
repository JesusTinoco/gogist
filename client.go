package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type response struct {
	status string
	header string
	body   string
}

func post(url, data, auth) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))

	if auth != null {
		req.SetBasicAuth(auth.username, auth.password)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)

	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return response{resp.Status, resp.Header, string(body)}
}
