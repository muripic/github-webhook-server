package issue

import (
	"database/sql"
	"fmt"
	"github-webhook-server/db"
)

const (
	issueTable        = "github.issue"
	labelTable        = "github.label"
	issueLabelTable   = "github.issue_label"
	issueCommentTable = "github.issue_comment"
)

var issueFields = []string{
	"id", "title", "number", "state", "body", "created_by",
	"url", "created_at", "updated_at", "closed_at",
}
var labelFields = []string{"id", "label"}
var issueLabelFields = []string{"issue_id", "label_id"}
var issueCommentFields = []string{
	"id", "issue_id", "body", "created_by", "created_at",
	"updated_at", "url",
}

/* Insert functions */

func insertIssue(i Issue) (sql.Result, error) {
	stmt := db.CreateInsertStatement(issueTable, issueFields)
	return DB.Exec(
		stmt, i.ID, i.Title, i.Number, i.State, i.Body, i.CreatedBy,
		i.URL, i.CreatedAt, i.UpdatedAt, i.ClosedAt,
	)
}

func insertLabel(l Label) (sql.Result, error) {
	stmt := db.CreateInsertStatement(labelTable, labelFields)
	return DB.Exec(stmt, l.ID, l.Label)
}

func insertIssueLabel(issueID int64, labelID int64) (sql.Result, error) {
	stmt := db.CreateInsertStatement(issueLabelTable, issueLabelFields)
	res, err := DB.Exec(stmt, issueID, labelID)
	if err != nil && !db.IsDuplicateKeyError(err) {
		panic(err)
	}
	return res, nil
}

func insertIssueComment(c IssueComment) (sql.Result, error) {
	stmt := db.CreateInsertStatement(issueCommentTable, issueCommentFields)
	return DB.Exec(
		stmt, c.ID, c.IssueID, c.Body, c.CreatedBy,
		c.CreatedAt, c.UpdatedAt, c.URL,
	)
}

/* Update functions */

func updateIssue(i Issue) (sql.Result, error) {
	stmt := db.CreateUpdateStatement(issueTable, issueFields[1:], i.ID)
	return DB.Exec(
		stmt, i.Title, i.Number, i.State, i.Body, i.CreatedBy,
		i.URL, i.CreatedAt, i.UpdatedAt, i.ClosedAt,
	)
}

func updateLabel(label Label) (sql.Result, error) {
	stmt := db.CreateUpdateStatement(labelTable, labelFields[1:], label.ID)
	return DB.Exec(stmt, label.Label)
}

func updateIssueComment(c IssueComment) (sql.Result, error) {
	stmt := db.CreateUpdateStatement(issueCommentTable, issueCommentFields[1:], c.ID)
	return DB.Exec(
		stmt, c.IssueID, c.Body, c.CreatedBy, c.CreatedAt, c.UpdatedAt, c.URL,
	)
}

/* Delete functions */

func deleteIssue(i Issue) {
	deleteIssueLabels(i)
	fmt.Println("Deleting issue...")
	db.Delete(DB, issueTable, "id", i.ID)
	fmt.Println("Issue deleted successfully.")
}

func deleteIssueLabels(i Issue) {
	fmt.Println("Deleting issue labels...")
	db.Delete(DB, issueLabelTable, "issue_id", i.ID)
	fmt.Println("Issue labels deleted successfully.")
}

func deleteIssueComment(c IssueComment) {
	fmt.Println("Deleting issue comment...")
	db.Delete(DB, issueCommentTable, "id", c.ID)
	fmt.Println("Issue comment deleted successfully.")
}

func deleteIssueComments(i Issue) {
	fmt.Println("Deleting all comments for issue...")
	db.Delete(DB, issueCommentTable, "issue_id", i.ID)
	fmt.Println("Issue comments deleted successfully.")
}

/* Insert or update functions */

func insertOrUpdateLabel(l Label) {
	// FIXME: refactor to remove repeated code
	_, err := insertLabel(l)
	if err == nil {
		fmt.Println("Label inserted successfully.")
		return
	}
	if db.IsDuplicateKeyError(err) {
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
		// TODO: Add delete logic. Use case: a label is deleted.
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
	if db.IsDuplicateKeyError(err) {
		fmt.Println("Issue already exists. Updating...")
		_, err := updateIssue(i)
		if err == nil {
			fmt.Println("Issue updated successfully.")
			return
		}
	}
	panic(err)
}

func insertOrUpdateIssueComment(c IssueComment) {
	// FIXME: refactor to remove repeated code
	_, err := insertIssueComment(c)
	if err == nil {
		fmt.Println("Issue comment inserted successfully.")
		return
	}
	if db.IsDuplicateKeyError(err) {
		fmt.Println("Issue comment already exists. Updating...")
		_, err := updateIssueComment(c)
		if err == nil {
			fmt.Println("Issue comment updated successfully.")
			return
		}
	}
	panic(err)
}
