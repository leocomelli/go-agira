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

	getIssue(client)
	getIssueEstimationForBoard(client)
}

func getIssue(client *jira.Client) {
	fmt.Println("GET ISSUE...")
	issue, _, err := client.Issues.Get(context.Background(), "816868", &jira.GetIssueOptions{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\t%s - %s\n", issue.ID, issue.Key)
}

func getIssueEstimationForBoard(client *jira.Client) {
	fmt.Println("GET ISSUE ESTIMATION...")
	issueEst, _, err := client.Issues.GetEstimationForBoard(context.Background(), "864189", 2881)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\t%s - %d\n", issueEst.FieldID, issueEst.Value)
}
