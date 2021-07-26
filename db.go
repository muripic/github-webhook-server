package main

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

var db *sql.DB
var dbConfig DBConfig

func loadDBConfig() {
	viper.UnmarshalKey("db", &dbConfig)
}

func getConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User,
		dbConfig.Password, dbConfig.DBName,
	)
}

func connectToDB() {
	var err error
	db, err = sql.Open("postgres", getConnectionString())
	if err != nil {
		panic(err)
	}
}

func checkDBConnection() {
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database!")
}

func createInsertStatement(table string, fields []string) string {
	var values []string
	for i := 1; i <= len(fields); i++ {
		values = append(values, fmt.Sprintf("$%d", i))
	}
	return fmt.Sprintf(
		"INSERT INTO %s(%s) VALUES (%s);",
		table,
		strings.Join(fields, ", "),
		strings.Join(values, ", "),
	)
}

func createUpdateStatement(table string, fields []string, id int64) string {
	var values []string
	for i, f := range fields {
		values = append(values, fmt.Sprintf("%s = $%d", f, i+1))
	}
	return fmt.Sprintf(
		"UPDATE %s SET %s WHERE id = %d",
		table, strings.Join(values, ", "), id,
	)
}

func isDuplicateKeyError(err error) bool {
	return strings.Contains(err.Error(), "duplicate key")
}
