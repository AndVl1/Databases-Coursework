package model

//easyjson:json
type Comment struct {
	Id             int    `json:"commentId"`
	Text           string `json:"text"`
	AttachmentPath string `json:"attachmentPath"`
	AuthorId       int    `json:"authorId"`
}

//easyjson:json
type Comments []*Comment
