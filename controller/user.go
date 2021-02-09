package controller

import (
	"context"
	"github.com/AndVl1/bugTrackerBackend/model"
	"github.com/AndVl1/bugTrackerBackend/storage"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func GetAllUsers(ctx echo.Context) error {
	users, _ := getRepoUsers()
	//if err != nil {
	//	return err
	//} TODO
	json, _ := users.MarshalJSON()

	return ctx.Blob(http.StatusOK, echo.MIMEApplicationJSON, json)
}

func GetUser(ctx echo.Context) error {
	id := ctx.Param("id")
	user, _ := getRepoUser(id)
	//if err != nil {
	//	return err
	//}
	return ctx.JSON(http.StatusOK, user)
}

func AddUser(ctx echo.Context) error {
	userJson := ctx.Param("user")
	log.Print(userJson)
	user := &model.User{}
	_ = user.UnmarshalJSON([]byte(userJson))
	_ = insertUser(user)
	return ctx.String(http.StatusOK, "OK")
}

func insertUser(user *model.User) error {
	pool := storage.GetDBInstance()
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v\n", err)
		return err
	}
	defer conn.Release()
	log.Print(user)
	row := conn.QueryRow(context.Background(),
		"INSERT INTO \"User\" (login, password, name, role) VALUES ($1, $2, $3, $4) RETURNING id",
		user.Login, user.Password, user.Name, user.Role)
	var id uint64
	err = row.Scan(&id)
	if err != nil {
		log.Printf("Unable to INSERT: %v\n", err)
		return err
	}
	return nil
}

func getRepoUser(id string) (model.User, error) {
	pool := storage.GetDBInstance()
	var user model.User
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v\n", err)
		return model.User{}, err
	}
	defer conn.Release()
	row := conn.QueryRow(context.Background(), "SELECT * FROM \"User\" WHERE userId=$1", id)
	_ = row.Scan(&user.Id, &user.Login, &user.Password, &user.Name, &user.Role)

	return user, nil
}

func getRepoUsers() (model.Users, error) {
	pool := storage.GetDBInstance()
	var users model.Users
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v\n", err)
		return nil, err
	}
	defer conn.Release()
	rows, _ := conn.Query(context.Background(), `SELECT * FROM "User"`)
	for rows.Next() {
		var user model.User
		_ = rows.Scan(&user.Id, &user.Login, &user.Password, &user.Name, &user.Role)
		log.Println(user.Name)
		users = append(users, &user)
	}

	return users, nil
}
