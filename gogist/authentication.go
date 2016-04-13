package gogist

import (
	"fmt"
	"gogist/client"
	"strings"
)

const url string = "https://api.github.com/authorizations"

func CreateNewToken(username string, password string) {
	code, _ := client.Post(url, `{"note": "Gogist2 tool", "scope": "gists"}`, username, password)

	if strings.Contains(code, "422") {
		panic("Token already exist. Remove it before continuing")
	} else if strings.Contains(code, "403") {
		panic("Failed to access to the URL")
	} else if strings.Contains(code, "201") {
		fmt.Println("Token created successfully. It will be safe at ~/gogist")
	}

}
