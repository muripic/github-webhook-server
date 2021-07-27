package push

import (
	"log"
	"strings"

	"github.com/google/go-github/github"
)

const targetRepo string = "muripic/content-gitops"
const targetBranch string = "master"

func changedFiles(e github.PushEvent) []string {
	var changedFiles []string
	for i := range e.Commits {
		changedFiles = append(changedFiles, e.Commits[i].Added...)
		changedFiles = append(changedFiles, e.Commits[i].Modified...)
		changedFiles = append(changedFiles, e.Commits[i].Removed...)
	}
	return changedFiles
}

func checkRepo(e github.PushEvent) bool {
	return *e.Repo.FullName == targetRepo
}

func checkBranch(e github.PushEvent) bool {
	branch := strings.Split(*e.Ref, "/")[2]
	return branch == targetBranch
}

func GetModifiedFiles(e github.PushEvent) {
	log.Print("Analyzing push event")
	if checkRepo(e) && checkBranch(e) {
		log.Printf("Getting modified files for %s repository", targetRepo)
		log.Printf("Files modified %s", changedFiles(e))
		return
	}
	log.Printf("Push does not match target repo or target branch")
}
