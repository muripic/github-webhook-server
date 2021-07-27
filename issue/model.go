package issue

import (
	"time"

	"github.com/google/go-github/github"
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
