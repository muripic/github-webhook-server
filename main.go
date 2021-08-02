package main

import (
	"fmt"
	"github-webhook-server/config"
	"github-webhook-server/issue"
	"github-webhook-server/push"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/go-github/github"
)

func parseWebHook(r *http.Request) (interface{}, error) {
	// Read payload into a []byte buffer
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading request body: %s", err)
	}
	defer r.Body.Close()
	// Parse webhook into event
	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		return nil, fmt.Errorf("error parsing webhook: %s", err)
	}
	return event, nil
}

func handleWebHook(w http.ResponseWriter, r *http.Request) {
	event, err := parseWebHook(r)
	if err != nil {
		log.Println(err)
	}
	// FIXME: remove switch, refactor function to return event
	// and call it in specific handlers
	switch e := event.(type) {
	case *github.IssuesEvent:
		handleIssueEvent(*e)
	case *github.IssueCommentEvent:
		handleIssueEvent(*e)
	case *github.PushEvent:
		handlePushEvent(*e)
	default:
		log.Printf("unknown event type %s\n", github.WebHookType(r))
	}
}

func handleIssueEvent(e interface{}) {
	issue.SaveIssueDataToDB(e)
}

func handlePushEvent(e github.PushEvent) {
	log.Print("Handling push event...")
	push.GetModifiedFiles(e)
}

func main() {
	log.Println("Server started")
	config.ReadConfig()
	// FIXME: each endpoint should have its own webhook handler
	http.HandleFunc("/issues", handleWebHook)
	http.HandleFunc("/pushes", handleWebHook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
