package main

import (
	"bytes"
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

	getEpic(client)
	listIssues(client)
	if write {
		partiallyUpdate(client)
	}
}

func getEpic(client *jira.Client) {
	fmt.Println("GET EPIC...")
	epic, _, err := client.Epics.Get(context.Background(), "523967")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\t%d - %s - %s - %s\n", epic.ID, epic.Name, epic.Key, epic.Color)
}

func listIssues(client *jira.Client) {
	fmt.Println("ISSUES...")

	issues, _, err := client.Epics.ListIssues(context.Background(), "523967", &jira.IssuesOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range issues {
		fmt.Printf("\tid: %s, key: %s, reporter: %s, status: %s\n",
			i.ID, i.Key, i.Fields.Reporter.DisplayName, i.Fields.Status.Name)
	}
}

func partiallyUpdate(client *jira.Client) {
	fmt.Println("PARTIALLY UPDATE...")

	epic := &jira.Epic{
		Name: "C.C.",
	}

	updatedEpic, resp, err := client.Epics.PartiallyUpdate(context.Background(), "523967", epic)
	if err != nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		newStr := buf.String()
		log.Println(resp.StatusCode)
		log.Println(newStr)
		log.Fatal(err)
	}

	fmt.Printf("epic: %d - %s - %s - %s\n", updatedEpic.ID, updatedEpic.Name, updatedEpic.Key, updatedEpic.Color)
}
