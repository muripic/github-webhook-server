package issue

import (
	"database/sql"
	"github-webhook-server/db"

	"github.com/google/go-github/github"
)

var DBCfg db.DBConfig
var DB *sql.DB

func SaveIssueToDB(e github.IssuesEvent) {
	DBCfg = db.LoadDBConfig()
	DB = db.ConnectToDB(DBCfg)
	db.CheckDBConnection(DB)
	defer DB.Close()
	issue := getIssue(e)
	if *e.Action == "deleted" {
		deleteIssue(issue)
		return
	}
	insertOrUpdateIssue(issue)
	insertOrUpdateLabels(issue)
}
