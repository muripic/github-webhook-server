package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/go-github/github"
	_ "github.com/lib/pq"
)

type Issue struct {
	ID        int64
	Title     string
	Number    int
	State     string
	Body      string
	CreatedBy string
	URL       string
	CreatedAt time.Time
	UpdatedAt time.Time
	ClosedAt  time.Time
	Labels    []Label
}

type Label struct {
	ID    int64
	Label string
}

/* Insert functions */

func insertIssue(i Issue) (sql.Result, error) {
	fields := GetStructFields(&i)
	fields = ArrayToSnakeCase(fields[:len(fields)-1])
	stmt := createInsertStatement("github.issue", fields)
	return db.Exec(
		stmt,
		i.ID, i.Title, i.Number, i.State,
		i.Body, i.CreatedBy, i.URL,
		i.CreatedAt, i.UpdatedAt, i.ClosedAt,
	)
}

func insertLabel(l Label) (sql.Result, error) {
	fields := ArrayToSnakeCase(GetStructFields(&l))
	stmt := createInsertStatement("github.label", fields)
	return db.Exec(stmt, l.ID, l.Label)
}

func insertIssueLabel(issueID int64, labelID int64) (sql.Result, error) {
	fields := []string{"issue_id", "label_id"}
	stmt := createInsertStatement("github.issue_label", fields)
	res, err := db.Exec(stmt, issueID, labelID)
	if err != nil && !isDuplicateKeyError(err) {
		panic(err)
	}
	return res, nil
}

/* Update functions */

func updateIssue(i Issue) (sql.Result, error) {
	fields := GetStructFields(&i)
	fields = ArrayToSnakeCase(fields[1 : len(fields)-1])
	stmt := createUpdateStatement("github.issue", fields, i.ID)
	return db.Exec(
		stmt, i.Title, i.Number, i.State,
		i.Body, i.CreatedBy, i.URL,
		i.CreatedAt, i.UpdatedAt, i.ClosedAt,
	)
}

func updateLabel(label Label) (sql.Result, error) {
	fields := []string{"label"}
	stmt := createUpdateStatement("github.label", fields, label.ID)
	return db.Exec(stmt, label.Label)
}

/* Delete functions */

func deleteIssue(issue Issue) {
	fmt.Println("Should delete issue.")
	// TODO: delete issue
	// Delete all labels with that issue id from issue_labels
}

/* Insert or update functions */

func insertOrUpdateLabel(l Label) {
	// FIXME: refactor to remove repeated code
	_, err := insertLabel(l)
	if err == nil {
		fmt.Println("Label inserted successfully.")
		return
	}
	if isDuplicateKeyError(err) {
		fmt.Println("Label already exists. Updating...")
		_, err := updateLabel(l)
		if err == nil {
			fmt.Println("Label updated successfully.")
			return
		}
	}
	panic(err)
}

func insertOrUpdateLabels(i Issue) {
	for _, l := range i.Labels {
		insertOrUpdateLabel(l)
		insertIssueLabel(i.ID, l.ID)
	}
}

func insertOrUpdateIssue(i Issue) {
	// FIXME: refactor to remove repeated code
	_, err := insertIssue(i)
	if err == nil {
		fmt.Println("Issue inserted successfully.")
		return
	}
	if isDuplicateKeyError(err) {
		fmt.Println("Issue already exists. Updating...")
		_, err := updateIssue(i)
		if err == nil {
			fmt.Println("Issue updated successfully.")
			return
		}
	}
	panic(err)
}

func getIssue(e github.IssuesEvent) Issue {
	var issue Issue
	issue.ID = *e.Issue.ID
	issue.Title = *e.Issue.Title
	issue.Number = *e.Issue.Number
	issue.State = *e.Issue.State
	issue.Body = *e.Issue.Body
	issue.CreatedBy = *e.Issue.User.Login
	issue.URL = *e.Issue.HTMLURL
	issue.CreatedAt = *e.Issue.CreatedAt
	issue.UpdatedAt = *e.Issue.UpdatedAt
	if *e.Action == "closed" {
		issue.ClosedAt = *e.Issue.ClosedAt
	}
	for _, l := range e.Issue.Labels {
		label := Label{*l.ID, *l.Name}
		issue.Labels = append(issue.Labels, label)
	}
	return issue
}

func saveIssueInfo(e github.IssuesEvent) {
	connectToDB()
	checkDBConnection()
	defer db.Close()
	issue := getIssue(e)
	if *e.Action == "deleted" {
		deleteIssue(issue)
		return
	}
	insertOrUpdateIssue(issue)
	insertOrUpdateLabels(issue)
}
