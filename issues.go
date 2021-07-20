package main

import "github.com/google/go-github/github"

func eventShouldBeSaved(e github.IssuesEvent) bool {
	return *e.Action == "opened" || *e.Action == "closed"
}

func getIssueInfo(e github.IssuesEvent) {
	title := *e.Issue.Title
	id := *e.Issue.ID
	number := *e.Issue.Number
	state := *e.Issue.State
	body := *e.Issue.Body
	createdAt := *e.Issue.CreatedAt
	updatedAt := *e.Issue.UpdatedAt
	user := *e.Issue.User.Login
	labels := e.Issue.Labels
	url := *e.Issue.HTMLURL
}
