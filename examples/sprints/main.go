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

	getSprint(client)
	if write {
		createSprint(client)
		updateSprint(client)
		partiallyUpdateSprint(client)
	}
}

func createSprint(client *jira.Client) {
	fmt.Println("CREATE SPRINT...")

	nSprint := &jira.NewSprint{
		Name:    "My First Sprint",
		BoardID: 2881,
	}

	sprint, _, err := client.Sprints.Create(context.Background(), nSprint)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\t%s - %d\n", sprint.Name, sprint.BoardID)
}

func getSprint(client *jira.Client) {
	fmt.Println("GET SPRINT...")
	sprint, _, err := client.Sprints.Get(context.Background(), 11392)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\t%d - %s - %d - %v\n", sprint.ID, sprint.Name, sprint.BoardID, sprint.Start)
}

func updateSprint(client *jira.Client) {
	fmt.Println("UPDATE SPRINT...")

	uSprint := &jira.Sprint{
		Name:    "My First Sprint Up1",
		Goal:    "I do not know",
		BoardID: 2881,
		State:   "future",
	}

	sprint, _, err := client.Sprints.Update(context.Background(), 11392, uSprint)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\t%d - %s - %d - %s - %v\n", sprint.ID, sprint.Name, sprint.BoardID, sprint.Goal, sprint.Start)
}

func partiallyUpdateSprint(client *jira.Client) {
	fmt.Println("PARTIALLY UPDATE SPRINT...")

	uSprint := &jira.Sprint{
		Name: "My First Sprint Up1 ***",
	}

	sprint, _, err := client.Sprints.PartiallyUpdate(context.Background(), 11392, uSprint)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\t%d - %s - %d - %s - %v\n", sprint.ID, sprint.Name, sprint.BoardID, sprint.Goal, sprint.Start)
}
