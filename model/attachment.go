package model

//easyjson:json
type Attachment struct {
	Id             int    `json:"attachmentId"`
	AuthorId       int    `json:"authorId"`
	AttachmentPath string `json:"attachmentPath"`
}

//easyjson:json
type Attachments []*Attachment
