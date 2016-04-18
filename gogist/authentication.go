package gogist

import (
	"encoding/json"
	"fmt"
	"gogist/client"
	"io/ioutil"
	"os"
	"strings"
)

const url string = "https://api.github.com/authorizations"

type TokenJSON struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
	//App            []App  `json:"app"`
	Token          string `json:"token"`
	HashedToken    string `json:"hashed_token"`
	TokenLastEight string `json:"token_last_eight"`
	Note           string `json:"note"`
	NoteUrl        string `json:"note_url"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	Scope          string `json:"scope"`
	Fingerprint    string `json:"fingerprint"`
}

type App struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	ClientId string `json:"client_id"`
}

var tokenPath string = os.Getenv("HOME") + "/.gogist"

func CreateNewToken(username string, password string) {

	code, body := client.Post(url, `{"note": "Gogist2 tool", "scope": "gists"}`, username, password)

	if strings.Contains(code, "422") {
		panic("Token already exist. Remove it before continuing")
	} else if strings.Contains(code, "403") {
		panic("Failed accessing to the URL")
	} else if strings.Contains(code, "201") {
		fmt.Println("Token created successfully. It will be safe at ~/gogist")
	}

	var tokenJson TokenJSON
	err := json.Unmarshal([]byte(body), &tokenJson)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(tokenPath, []byte(tokenJson.Token), 0660)
	if err != nil {
		panic(err)
	}

	fmt.Println("Token saved successfully in ", tokenPath)
}
