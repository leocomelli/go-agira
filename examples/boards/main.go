package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

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
		createBoard(client)
		deleteBoard(client)
	}
	listBoards(client)
	getBoard(client)
	getBoardConfiguration(client)
	listEpics(client)
	listBacklogIssues(client)
	listSprints(client)
	listIssues(client)
	listIssuesForEpic(client)
	listIssuesWithoutEpic(client)
	listProjects(client)
	listIssuesForSprint(client)
	listVersions(client)
}

func createBoard(client *jira.Client) {
	fmt.Println("CREATE...")

	newBoard := &jira.NewBoard{Name: "BOARDTEST", Type: "scrum", FilterID: 40079}

	b, resp, err := client.Boards.CreateBoard(context.Background(), newBoard)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(resp.StatusCode)
	fmt.Println(b.ID)
}

func deleteBoard(client *jira.Client) {
	fmt.Println("DELETE...")

	resp, err := client.Boards.DeleteBoard(context.Background(), 5597)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(resp.StatusCode)
}

func listBoards(client *jira.Client) {
	fmt.Println("LIST BOARD...")

	opts := &jira.BoardsOptions{
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

func getBoardConfiguration(client *jira.Client) {
	fmt.Println("GET BOARD CONFIGURATION...")
	configuration, _, err := client.Boards.GetBoardConfiguration(context.Background(), 2881)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\t%d - %s - %s\n", configuration.ID, configuration.Name, configuration.SelfLink)
}

func listEpics(client *jira.Client) {
	fmt.Println("LIST EPICS...")

	epics, _, err := client.Boards.ListEpics(context.Background(), 2881, &jira.EpicsOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range epics {
		fmt.Printf("\tid: %d, name: %s, done: %v\n", e.ID, e.Name, e.Done)
	}
}

func listBacklogIssues(client *jira.Client) {
	fmt.Println("BACKLOG...")

	issues, _, err := client.Boards.ListBacklogIssues(context.Background(), 2881, &jira.IssuesOptions{})
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

	sprints, _, err := client.Boards.ListSprints(context.Background(), 2881, &jira.SprintsOptions{})
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

	issues, _, err := client.Boards.ListIssues(context.Background(), 2881, &jira.IssuesOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range issues {
		fmt.Printf("\tid: %s, key: %s, reporter: %s, status: %s\n",
			i.ID, i.Key, i.Fields.Reporter.DisplayName, i.Fields.Status.Name)
	}
}

func listIssuesForEpic(client *jira.Client) {
	fmt.Println("ISSUES FOR EPIC...")

	issues, _, err := client.Boards.ListIssuesForEpic(context.Background(), 2881, 523967, &jira.IssuesOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range issues {
		fmt.Printf("\tid: %s, key: %s, reporter: %s, status: %s, epic: %s\n",
			i.ID, i.Key, i.Fields.Reporter.DisplayName, i.Fields.Status.Name, i.Fields.Epic.Name)
	}
}

func listIssuesWithoutEpic(client *jira.Client) {
	fmt.Println("ISSUES WITHOUT EPIC...")

	issues, _, err := client.Boards.ListIssuesWithoutEpic(context.Background(), 2881, &jira.IssuesOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range issues {
		fmt.Printf("\tid: %s, key: %s, reporter: %s, status: %s, epic: %v\n",
			i.ID, i.Key, i.Fields.Reporter.DisplayName, i.Fields.Status.Name, i.Fields.Epic)
	}
}

func listProjects(client *jira.Client) {
	fmt.Println("PROJECTS...")

	projects, _, err := client.Boards.ListProjects(context.Background(), 2881, &jira.ProjectsOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range projects {
		fmt.Printf("\tid: %s, key: %s, category: %v\n",
			p.ID, p.Key, p.Category)
	}
}

func listIssuesForSprint(client *jira.Client) {
	fmt.Println("ISSUES FOR SPRINT...")

	issues, _, err := client.Boards.ListIssuesForSprint(context.Background(), 2881, 10224, &jira.IssuesOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range issues {
		fmt.Printf("\tid: %s, key: %s, reporter: %s\n",
			i.ID, i.Key, i.Fields.Reporter.DisplayName)
	}
}

func listVersions(client *jira.Client) {
	fmt.Println("VERSIONS...")

	versions, _, err := client.Boards.ListVersions(context.Background(), 2881, &jira.VersionsOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range versions {
		fmt.Printf("\tid: %d, name: %s, released: %v, projectId: %v\n",
			v.ID, v.Name, v.Released, v.ProjectID)
	}
}
