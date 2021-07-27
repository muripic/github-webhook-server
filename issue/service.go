package issue

import (
	"database/sql"
	"github-webhook-server/db"

	"github.com/google/go-github/github"
)

var DBCfg db.DBConfig
var DB *sql.DB

func setUpDB() {
	db.LoadDBConfig(&DBCfg)
	DB = db.ConnectToDB(DBCfg)
}

func SaveIssueToDB(e github.IssuesEvent) {
	setUpDB()
	defer DB.Close()
	i := getIssue(e)
	if *e.Action == "deleted" {
		deleteIssue(i)
		return
	}
	insertOrUpdateIssue(i)
	insertOrUpdateLabels(i)
}

func SaveIssueCommentToDB(e github.IssueCommentEvent) {
	setUpDB()
	defer DB.Close()
	c := getIssueComment(e)
	if *e.Action == "deleted" {
		deleteIssueComment(c)
		return
	}
	insertOrUpdateIssueComment(c)
}
