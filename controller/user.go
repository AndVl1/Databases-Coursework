package controller

import (
	"context"
	"errors"
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

func LoginUser(ctx echo.Context) error {
	log.Printf("email(%s) pw(%s)", ctx.FormValue("email"), ctx.FormValue("password"))
	user, err := login(ctx.FormValue("email"), ctx.FormValue("password"))
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err)
	}
	userJson, _ := user.MarshalJSON()
	return ctx.Blob(http.StatusOK, echo.MIMEApplicationJSON, userJson)
}

func GetUser(ctx echo.Context) error {
	id := ctx.Param("id")
	user, _ := getRepoUser(id)
	json, _ := user.MarshalJSON()
	//if err != nil {
	//	return err
	//}
	return ctx.Blob(http.StatusOK, echo.MIMEApplicationJSON, json)
}

//func AddUser(ctx echo.Context) error {
//	userJson := ctx.FormValue("user")
//	log.Print(userJson)
//	user := &model.User{}
//	_ = user.UnmarshalJSON([]byte(userJson))
//	_, _ = insertUser(user)
//	return ctx.String(http.StatusOK, "OK")
//}

func CheckUser(ctx echo.Context) error {
	emailJson := ctx.FormValue("email")
	registered, err := checkRegisteredUser(emailJson)
	if err != nil {
		log.Println(err)
		return err
	}
	if registered {
		return ctx.JSON(http.StatusBadRequest, "")
	} else {
		return ctx.JSON(http.StatusOK, "")
	}
}

func RegisterUser(ctx echo.Context) error {
	user := &model.User{
		Id:       0,
		Login:    ctx.FormValue("email"),
		Password: ctx.FormValue("password"),
		Name:     ctx.FormValue("name"),
	}
	log.Println(user.Login)
	id, err := insertUser(user)
	if err != nil {
		log.Println(err)
		return err
	}
	user.Id = id
	var userJson, _ = user.MarshalJSON()
	return ctx.Blob(http.StatusOK, echo.MIMEApplicationJSON, userJson)
}

func login(login string, password string) (*model.User, error) {
	pool := storage.GetDBInstance()
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v\n", err)
		return nil, err
	}
	defer conn.Release()
	var user model.User
	row := conn.QueryRow(context.Background(), "SELECT * FROM \"User\" WHERE login=$1", login)
	_ = row.Scan(&user.Id, &user.Login, &user.Password, &user.Name)
	if user.Password != password {
		return nil, errors.New("wrong password")
	} else {
		return &user, nil
	}
}

func insertUser(user *model.User) (uint64, error) {
	pool := storage.GetDBInstance()
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v\n", err)
		return 0, err
	}
	defer conn.Release()
	log.Print(user)
	row := conn.QueryRow(context.Background(),
		"INSERT INTO \"User\" (login, password, name) VALUES ($1, $2, $3) RETURNING userId",
		user.Login, user.Password, user.Name)
	var id uint64
	err = row.Scan(&id)
	if err != nil {
		log.Printf("Unable to INSERT: %v\n", err)
		return 0, err
	}
	return id, nil
}

func checkRegisteredUser(login string) (bool, error) {
	pool := storage.GetDBInstance()
	var count int
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v\n", err)
		return true, err
	}
	defer conn.Release()
	row := conn.QueryRow(context.Background(), "SELECT COUNT(*) FROM \"User\" WHERE login=$1", login)
	_ = row.Scan(&count)
	return count > 0, nil
}

func getRepoUser(id string) (*model.User, error) {
	pool := storage.GetDBInstance()
	var user model.User
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v\n", err)
		return nil, err
	}
	defer conn.Release()
	row := conn.QueryRow(context.Background(), "SELECT * FROM \"User\" WHERE userId=$1", id)
	_ = row.Scan(&user.Id, &user.Login, &user.Password, &user.Name)

	return &user, nil
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
		_ = rows.Scan(&user.Id, &user.Login, &user.Password, &user.Name)
		log.Println(user.Name)
		users = append(users, &user)
	}

	return users, nil
}
