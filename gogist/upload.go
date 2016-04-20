package gogist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gogist/client"
	"io/ioutil"
	"path"
	"strings"
)

type Gist struct {
	Description string          `json:"description"`
	Public      bool            `json:"public"`
	Files       map[string]File `json:"files"`
}

type File struct {
	Content string `json:"content"`
}

type Response struct {
	Url string `json:"url"`
}

func CreateGist(fileList []string, isPublic bool, description string, accessToken string) {

	files := make(map[string]File)

	for i := 0; i < len(fileList); i++ {
		content, err := ioutil.ReadFile(fileList[i])
		if err != nil {
			panic(err)

		}
		fileName := path.Base(fileList[i])
		files[fileName] = File{Content: string(content)}

	}

	gist := Gist{Description: description, Public: isPublic, Files: files}

	buf, err := json.Marshal(gist)
	if err != nil {
		panic(err)

	}
	jsonBody := bytes.NewBuffer(buf)

	var newUrl string = "https://api.github.com/gists?access_token=" + accessToken

	code, _ := client.Post(newUrl, jsonBody.String(), "", "")

	if strings.Contains(code, "201") {
		fmt.Println("Gist create successfully")
	}

}
