package gogist

import (
	"fmt"
	"gogist/client"
)

const url string = "https://api.github.com/authorizations"

func CreateNewToken(username string, password string) {
	code, body := client.Post(url, `{"note": "Gogist2 tool", "scope": "gists"}`, username, password)
	fmt.Println(code)
	fmt.Println(body)
}
