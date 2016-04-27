package client

import (
	"io"
	"io/ioutil"
	"net/http"
)

func Post(url string, data io.Reader, username string, password string) (string, string) {
	req, err := http.NewRequest("POST", url, data)

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
	return resp.Status, string(body)
}
