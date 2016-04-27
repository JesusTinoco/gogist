package gogist

import (
	"os"
)

var tokenPath string = os.Getenv("HOME") + "/.gogist"
var baseUrl string = "https://api.github.com/"
var authorizationsUrl string = baseUrl + "authorizations"
var gistsUrl string = baseUrl + "gists?access_token="
