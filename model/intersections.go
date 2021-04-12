package model

//easyjson:json
type ProjectUser struct {
	UserId    uint64 `json:"userId"`
	ProjectId uint64 `json:"projectId"`
}

//easyjson:json
type ProjectUsers []*ProjectUser

//easyjson:json
type IssueRecursive struct {
	IssueId         uint64 `json:"issueId"`
	ContainsIssueId uint64 `json:"containsIssueId"`
}

//easyjson:json
type IssuesRecursive []*IssueRecursive
