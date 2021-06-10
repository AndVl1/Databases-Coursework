package model

//easyjson:json
type Comment struct {
	Id       int    `json:"commentId"`
	Text     string `json:"text"`
	AuthorId int    `json:"authorId"`
	IssueId  int    `json:"issueId"`
	Date     int64  `json:"date"`
}

//easyjson:json
type Comments []*Comment

//easyjson:json
type CommentRest struct {
	Id      int    `json:"commentId"`
	Text    string `json:"text"`
	Author  User   `json:"author"`
	Date    int64  `json:"date"`
	IssueId int    `json:"issueId"`
}

//easyjson:json
type CommentsRest []*CommentRest
