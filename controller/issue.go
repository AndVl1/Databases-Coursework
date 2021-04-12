package controller

import (
	"context"
	"github.com/AndVl1/bugTrackerBackend/model"
	"github.com/AndVl1/bugTrackerBackend/storage"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func GetAllIssues(ctx echo.Context) error {
	bugs, _ := getRepoIssues()
	//if err != nil {
	//	return err
	//} TODO
	json, _ := bugs.MarshalJSON()

	return ctx.Blob(http.StatusOK, echo.MIMEApplicationJSON, json)
}

func GetIssue(ctx echo.Context) error {
	id := ctx.Param("id")
	bug, _ := getRepoIssue(id)
	//if err != nil {
	//	return err
	//}
	json, _ := bug.MarshalJSON()
	return ctx.Blob(http.StatusOK, echo.MIMEApplicationJSON, json)
}

func AddIssue(ctx echo.Context) error {
	bugJson := ctx.Param("issue")
	log.Print(bugJson)
	issue := &model.Issue{}
	_ = issue.UnmarshalJSON([]byte(bugJson))
	_ = insertIssue(issue)
	return ctx.String(http.StatusOK, "OK")
}

func insertIssue(bug *model.Issue) error {
	pool := storage.GetDBInstance()
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v\n", err)
		return err
	}
	defer conn.Release()
	log.Print(bug)
	row := conn.QueryRow(context.Background(),
		"INSERT INTO Issue (name, description, status, authorId) VALUES ($1, $2, $3, $4) RETURNING id",
		bug.Name, bug.Description, bug.Status, bug.AuthorId)
	var id uint64
	err = row.Scan(&id)
	if err != nil {
		log.Printf("Unable to INSERT: %v\n", err)
		return err
	}
	return nil
}

func getRepoIssue(id string) (model.Issue, error) {
	pool := storage.GetDBInstance()
	var bug model.Issue
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v\n", err)
		return model.Issue{}, err
	}
	defer conn.Release()
	row := conn.QueryRow(context.Background(), "SELECT * FROM Issue WHERE bugId=$1", id)
	_ = row.Scan(&bug.Id, &bug.Name, &bug.Description, &bug.Status, &bug.AuthorId)

	return bug, nil
}

func getRepoIssues() (model.Issues, error) {
	pool := storage.GetDBInstance()
	var bugs model.Issues
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v\n", err)
		return nil, err
	}
	defer conn.Release()
	rows, _ := conn.Query(context.Background(), `SELECT * FROM Issues`)
	for rows.Next() {
		var bug model.Issue
		_ = rows.Scan(&bug.Id, &bug.Name, &bug.Description, &bug.Status)
		log.Println(bug.Name)
		bugs = append(bugs, &bug)
	}

	return bugs, nil
}
