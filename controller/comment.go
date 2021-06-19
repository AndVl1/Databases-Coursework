package controller

import (
	"context"
	"github.com/AndVl1/bugTrackerBackend/model"
	"github.com/AndVl1/bugTrackerBackend/storage"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

func AddComment(ctx echo.Context) error {
	authorId, _ := strconv.Atoi(ctx.FormValue("authorId"))
	issueId, _ := strconv.Atoi(ctx.Param("id"))
	date, _ := strconv.ParseInt(ctx.FormValue("date"), 10, 64)

	comment := model.Comment{
		Id:       0,
		Text:     ctx.FormValue("text"),
		AuthorId: authorId,
		IssueId:  issueId,
		Date:     date,
	}
	var err error
	comment.Id, err = addCommentRepo(comment)
	if err != nil {
		return ctx.String(http.StatusNotImplemented, err.Error())
	}
	commentJson, err := comment.MarshalJSON()
	if err != nil {
		return ctx.String(http.StatusNotImplemented, err.Error())
	}
	return ctx.Blob(http.StatusOK, echo.MIMEApplicationJSON, commentJson)
}

func GetComments(ctx echo.Context) error {
	comments, _ := getCommentsRepo(ctx.Param("id"))
	json, _ := comments.MarshalJSON()
	log.Println(string(json))
	return ctx.Blob(http.StatusOK, echo.MIMEApplicationJSON, json)
}

func getCommentsRepo(issueId string) (model.CommentsRest, error) {
	pool := storage.GetDBInstance()
	var comments model.CommentsRest
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v\n", err)
		return nil, err
	}
	defer conn.Release()
	rows, _ := conn.Query(context.Background(),
		"SELECT commentId, commentText, commentDate, issueId, userId, name, login "+
			"FROM CommentView WHERE issueId=$1 ORDER BY commentId", issueId)
	for rows.Next() {
		var comment model.CommentRest
		_ = rows.Scan(&comment.Id, &comment.Text, &comment.Date,
			&comment.IssueId, &comment.Author.Id, &comment.Author.Name, &comment.Author.Login)
		comments = append(comments, &comment)
	}

	return comments, nil
}

func addCommentRepo(comment model.Comment) (int, error) {
	pool := storage.GetDBInstance()
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v\n", err)
		return 0, err
	}
	defer conn.Release()
	projectRow := conn.QueryRow(
		context.Background(),
		"INSERT INTO Comment(commentText, commentDate, authorId, issueId) "+
			"VALUES ($1, $2, $3, $4) RETURNING commentId",
		comment.Text, comment.Date, comment.AuthorId, comment.IssueId)
	err = projectRow.Scan(&comment.Id)
	if err != nil {
		log.Printf("Unable to INSERT: %v\n", err)
		return 0, err
	}
	return comment.Id, nil
}
