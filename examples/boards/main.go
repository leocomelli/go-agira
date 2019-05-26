package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/leocomelli/jira"
)

var url, user, pass string

func init() {
	url = os.Getenv("JIRA_URL")
	user = os.Getenv("JIRA_USER")
	pass = os.Getenv("JIRA_PASS")
}

func main() {

	tp := &jira.BasicAuthTransport{
		Username: user,
		Password: pass,
	}
	client, err := jira.NewClient(url, tp.Client())
	if err != nil {
		log.Fatal(err)
	}

	listBoards(client)
	getBoard(client)
	listEpics(client)
	listBacklogIssues(client)
	listSprints(client)
	listIssues(client)
}

func listBoards(client *jira.Client) {
	fmt.Println("LIST BOARD...")

	opts := &jira.ListBoardsOptions{
		ProjectKeyOrID: "CBD",
		MaxResults:     1,
	}

	boards, resp, err := client.Boards.ListBoards(context.Background(), opts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\tmax results: %d\n", resp.MaxResults)
	fmt.Printf("\tstart at: %d\n", resp.StartAt)
	fmt.Printf("\tis last: %v\n", resp.IsLast)

	for _, b := range boards {
		fmt.Printf("\t%d - %s\n", b.ID, b.Name)
	}
}

func getBoard(client *jira.Client) {
	fmt.Println("GET BOARD...")
	board, _, err := client.Boards.GetBoard(context.Background(), 2881)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\t%d - %s - %s\n", board.ID, board.Name, board.Type)
}

func listEpics(client *jira.Client) {
	fmt.Println("LIST EPICS...")

	epics, _, err := client.Boards.ListEpics(context.Background(), 2881, &jira.ListEpicsOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range epics {
		fmt.Printf("\tid: %d, name: %s, done: %v\n", e.ID, e.Name, e.Done)
	}
}

func listBacklogIssues(client *jira.Client) {
	fmt.Println("BACKLOG...")

	issues, _, err := client.Boards.ListBacklogIssues(context.Background(), 2881, &jira.ListIssuesOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range issues {
		fmt.Printf("\tid: %s, key: %s, reporter: %s, status: %s\n",
			i.ID, i.Key, i.Fields.Reporter.DisplayName, i.Fields.Status.Name)
	}
}

func listSprints(client *jira.Client) {
	fmt.Println("LIST SPRINTS...")

	sprints, _, err := client.Boards.ListSprints(context.Background(), 2881, &jira.ListSprintsOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range sprints {
		fmt.Printf("\tid: %d, name: %s, state: %s, start: %v, end: %v\n",
			s.ID, s.Name, s.State, s.Start.Format(time.RFC3339Nano), s.End.Format(time.RFC3339Nano))
	}
}

func listIssues(client *jira.Client) {
	fmt.Println("ISSUES...")

	issues, _, err := client.Boards.ListIssues(context.Background(), 2881, &jira.ListIssuesOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range issues {
		fmt.Printf("\tid: %s, key: %s, reporter: %s, status: %s\n",
			i.ID, i.Key, i.Fields.Reporter.DisplayName, i.Fields.Status.Name)
	}
}
