package storage

import (
	"context"
	"log"

	"github.com/AndVl1/bugTrackerBackend/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func NewDB(params ...string) *pgxpool.Pool {
	var err error
	conString := config.GetPostgresConnectionString()

	log.Print(conString)

	DB, err = pgxpool.Connect(context.Background(), conString)

	if err != nil {
		log.Panic(err)
	}

	return DB
}

func GetDBInstance() *pgxpool.Pool {
	return DB
}
