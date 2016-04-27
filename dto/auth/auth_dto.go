package AuthDto

type Response struct {
	Id          int      `json:"id"`
	Url         string   `json:"url"`
	Scope       []string `json:"scope"`
	Token       string   `json:"token"`
	HashedToken string   `json:"hashed_token"`
	//App            []App    `json:"app"`
	TokenLastEight string `json:"token_last_eight"`
	Note           string `json:"note"`
	NoteUrl        string `json:"note_url"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	Fingerprint    string `json:"fingerprint"`
}

type App struct {
	Url      string `json:"url"`
	Name     string `json:"name"`
	ClientId string `json:"client_id"`
}
