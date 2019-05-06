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
	}

	boards, resp, err := client.Boards.ListBoards(context.Background(), opts)

	fmt.Printf("max results: %d\n", resp.MaxResults)
	fmt.Printf("start at: %d\n", resp.StartAt)
	fmt.Printf("is last: %v\n", resp.IsLast)
	fmt.Println("boards: ")

	for _, b := range boards {
		fmt.Printf("\t %d - %v\n", b.ID, b.Name)
	}

}
