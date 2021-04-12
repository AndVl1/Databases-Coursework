package model

//easyjson:json
type Issue struct {
	Id                 int    `json:"bugId"`
	Name               string `json:"name"`
	ProjectIssueNumber int    `json:"projectIssueNumber"`
	Description        string `json:"description"`
	Status             string `json:"status"`
	AuthorId           int    `json:"authorId"`
}

//easyjson:json
type Issues []*Issue
