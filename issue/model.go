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

type IssueComment struct {
	ID        int64
	IssueID   int64
	Body      string
	CreatedBy string
	CreatedAt time.Time
	UpdatedAt time.Time
	URL       string
}

func getIssue(e github.IssuesEvent) Issue {
	var i Issue
	i.ID = *e.Issue.ID
	i.Title = *e.Issue.Title
	i.Number = *e.Issue.Number
	i.State = *e.Issue.State
	i.Body = *e.Issue.Body
	i.CreatedBy = *e.Issue.User.Login
	i.URL = *e.Issue.HTMLURL
	i.CreatedAt = *e.Issue.CreatedAt
	i.UpdatedAt = *e.Issue.UpdatedAt
	if *e.Action == "closed" {
		i.ClosedAt = *e.Issue.ClosedAt
	}
	for _, l := range e.Issue.Labels {
		label := Label{*l.ID, *l.Name}
		i.Labels = append(i.Labels, label)
	}
	return i
}

func getIssueComment(e github.IssueCommentEvent) IssueComment {
	var c IssueComment
	c.ID = *e.Comment.ID
	c.IssueID = *e.Issue.ID
	c.Body = *e.Comment.Body
	c.CreatedBy = *e.Comment.User.Login
	c.CreatedAt = *e.Comment.CreatedAt
	c.UpdatedAt = *e.Comment.UpdatedAt
	c.URL = *e.Comment.HTMLURL
	return c
}
