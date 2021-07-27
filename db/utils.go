package db

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func LoadDBConfig() DBConfig {
	var cfg DBConfig
	viper.UnmarshalKey("db", &cfg)
	return cfg
}

func GetConnectionString(cfg DBConfig) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName,
	)
}

func ConnectToDB(cfg DBConfig) *sql.DB {
	DB, err := sql.Open("postgres", GetConnectionString(cfg))
	if err != nil {
		panic(err)
	}
	return DB
}

func CheckDBConnection(DB *sql.DB) {
	err := DB.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database!")
}

func CreateInsertStatement(table string, fields []string) string {
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

func CreateUpdateStatement(table string, fields []string, id int64) string {
	var values []string
	for i, f := range fields {
		values = append(values, fmt.Sprintf("%s = $%d", f, i+1))
	}
	return fmt.Sprintf(
		"UPDATE %s SET %s WHERE id = %d",
		table, strings.Join(values, ", "), id,
	)
}

func IsDuplicateKeyError(err error) bool {
	return strings.Contains(err.Error(), "duplicate key")
}
