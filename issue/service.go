package issue

import (
	"database/sql"
	"github-webhook-server/db"
	"log"

	"github.com/google/go-github/github"
)

var DBCfg db.DBConfig
var DB *sql.DB

func setUpDB() {
	db.LoadDBConfig(&DBCfg)
	DB = db.ConnectToDB(DBCfg)
}

func saveIssueToDB(e github.IssuesEvent) {
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

func saveIssueCommentToDB(e github.IssueCommentEvent) {
	setUpDB()
	defer DB.Close()
	c := getIssueComment(e)
	if *e.Action == "deleted" {
		deleteIssueComment(c)
		return
	}
	insertOrUpdateIssueComment(c)
}

func SaveIssueDataToDB(event interface{}) {
	switch e := event.(type) {
	case github.IssuesEvent:
		log.Print("Saving issue info to database...")
		saveIssueToDB(e)
	case github.IssueCommentEvent:
		log.Print("Saving issue comment info to database...")
		saveIssueCommentToDB(e)
	default:
		log.Print("Event type must be issue or issue comment")
	}
}
