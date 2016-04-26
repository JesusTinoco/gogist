package Dto

type Response struct {
	Url         string                 `json:"url"`
	ForksUrl    string                 `json:"forks_url"`
	CommitsUrl  string                 `json:"commits_url"`
	Id          string                 `json:"id"`
	Description string                 `json:"description"`
	Public      bool                   `json:"public"`
	Owner       Owner                  `json:"owner"`
	User        User                   `json:"user"`
	Files       map[string]FileDetails `json:"files"`
	Truncated   bool                   `json:"truncated"`
	Comments    int                    `json:"comments"`
	CommentsUrl string                 `json:"comments_url"`
	HtmlUrl     string                 `json:"html_url"`
	GitPullUrl  string                 `json:"git_pull_url"`
	GitPushUrl  string                 `json:"git_push_url"`
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
	Forks       []Fork                 `json:"forks"`
	History     []History              `json:"history"`
}

type Owner struct {
}

type User struct {
}

type FileDetails struct {
}

type Fork struct {
}

type History struct {
}

type Gist struct {
	Description string          `json:"description"`
	Public      bool            `json:"public"`
	Files       map[string]File `json:"files"`
}

type File struct {
	Content string `json:"content"`
}
