package gogist

import (
	"encoding/json"
	"fmt"
	"gogist/client"
	"gogist/gogist/dto"
	"io/ioutil"
	"os"
	"strings"
)

const url string = "https://api.github.com/authorizations"

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
