package gogist

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"gogist/client"
	"gogist/dto/auth"
	"io/ioutil"
	"os"
	"strings"
)

func CreateNewToken() {

	username, password := askForCredentials()

	code, body := client.Post(authorizationsUrl, bytes.NewBufferString(`{"note": "Gogist", "scopes": ["gist"]}`), strings.Trim(username, "\n"), strings.Trim(password, "\n"))

	if strings.Contains(code, "422") {
		panic("Token already exist. Remove it before continuing")
	} else if strings.Contains(code, "403") {
		panic("Failed accessing to the URL")
	} else if strings.Contains(code, "401") {
		panic("Bad credentials")
	} else if strings.Contains(code, "201") {
		fmt.Println("Token created successfully. It will be safe at ~/gogist")
	} else {
		panic(code + " " + body)
	}

	var tokenJson AuthDto.Response
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

func askForCredentials() (string, string) {
	fmt.Println("Obtaining OAuth2 access token from GitHub.")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter GitHub username: ")
	username, _ := reader.ReadString('\n')
	fmt.Print("Enter Github password: ")
	password, _ := reader.ReadString('\n')
	return strings.Trim(username, "\n"), strings.Trim(password, "\n")
}
