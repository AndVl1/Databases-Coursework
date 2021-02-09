package model

//easyjson:json
type Bug struct {
	Id          int    `json:"bugId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	AuthorId    int    `json:"authorId"`
}

//easyjson:json
type Bugs []*Bug
