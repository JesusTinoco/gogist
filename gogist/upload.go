package gogist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gogist/client"
	"gogist/dto/gists"
	"io/ioutil"
	"path"
	"strings"
)

func CreateGist(fileList []string, filename string, isPublic bool, description string) {

	accessToken, err := ioutil.ReadFile(tokenPath)
	if err != nil {
		panic(err)
	}

	files := make(map[string]GistsDto.File)

	for i := 0; i < len(fileList); i++ {
		content, err := ioutil.ReadFile(fileList[i])
		if err != nil {
			panic(err)

		}
		fileName := path.Base(fileList[i])
		files[fileName] = GistsDto.File{Content: string(content)}
	}

	gist := GistsDto.Gist{Description: description, Public: isPublic, Files: files}

	buf, err := json.Marshal(gist)
	if err != nil {
		panic(err)

	}
	jsonBody := bytes.NewBuffer(buf)
	var newUrl string = gistsUrl + strings.Trim(string(accessToken), " ")
	code, body := client.Post(newUrl, jsonBody, "", "")

	if strings.Contains(code, "201") {
		fmt.Println("Gist create successfully")
	} else {
		panic(code + body)
	}

}
