package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/leocomelli/jira"
)

func main() {

	url := os.Getenv("JIRA_URL")
	user := os.Getenv("JIRA_USER")
	pass := os.Getenv("JIRA_PASS")

	tp := &jira.BasicAuthTransport{
		Username: user,
		Password: pass,
	}
	client, err := jira.NewClient(url, tp.Client())
	if err != nil {
		log.Fatal(err)
	}

	opts := &jira.ListBoardsOptions{
		ProjectKeyOrID: "CBD",
		MaxResults:     1,
	}

	boards, resp, err := client.Boards.ListBoards(context.Background(), opts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("max results: %d\n", resp.MaxResults)
	fmt.Printf("start at: %d\n", resp.StartAt)
	fmt.Printf("is last: %v\n", resp.IsLast)
	fmt.Println("boards: ")

	for _, b := range boards {
		fmt.Printf("\t %d - %v\n", b.ID, b.Name)

		sprints, _, err := client.Boards.ListSprints(context.Background(), b.ID, nil)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("sprints: ")
		for _, s := range sprints {
			fmt.Printf("\t id: %d, name: %s, state: %s, start: %v, end: %v, complete: %v\n", s.ID, s.Name, s.State, s.Start, s.End, s.Complete)
		}

		epics, _, err := client.Boards.ListEpics(context.Background(), b.ID, &jira.ListEpicsOptions{})

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("epics: ")
		for _, e := range epics {
			fmt.Printf("\t id: %d, name: %s, done: %v\n", e.ID, e.Name, e.Done)
		}
	}

}
