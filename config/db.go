package config

import (
	"fmt"
	"os"
)

var (
	DBUser     = os.Getenv("DBUser")
	DBPassword = os.Getenv("DBPassword")
	DBName     = os.Getenv("DBName")
	DBHost     = os.Getenv("DBHost")
	DBPort     = os.Getenv("DBPort")
	DBType     = "postgres"
)

//var (
//	DBUser     = "bugtrack"
//	DBPassword = "bugtrack_pass"
//	DBName     = "bugtrackcourse"
//	DBHost     = "0.0.0.0"
//	DBPort     = "5432"
//	DBType     = "postgres"
//)

func GetDBType() string {
	return DBType
}

func GetPostgresConnectionString() string {
	dataBase := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		DBHost,
		DBPort,
		DBUser,
		DBName,
		DBPassword)
	return dataBase
}
