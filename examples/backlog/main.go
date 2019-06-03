package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/leocomelli/jira"
)

var (
	url, user, pass string
	write           bool
	err             error
)

func init() {
	url = os.Getenv("JIRA_URL")
	user = os.Getenv("JIRA_USER")
	pass = os.Getenv("JIRA_PASS")

	writeStr := os.Getenv("JIRA_WRITE_SRV")
	if writeStr == "" {
		writeStr = "false"
	}

	write, err = strconv.ParseBool(writeStr)
	if err != nil {
		log.Fatal(err)
	}
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

	if write {
		moveIssuesTo(client)
	}
}

func moveIssuesTo(client *jira.Client) {
	fmt.Println("MOVE ISSUES TO...")

	keys := &jira.IssueKeys{
		Issues: []string{"MCP-637"},
	}

	ok, _, err := client.Backlog.MoveIssuesTo(context.Background(), keys)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\t%v\n", ok)
}
