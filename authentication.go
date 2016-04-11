package authentication

import "client"

const url string = "https://gist.github.com/authorizations"

func createNewToken(auth) {
	client.post(url, `{"note": "Gogist tool", "scope": "gists"}`, auth)
}
