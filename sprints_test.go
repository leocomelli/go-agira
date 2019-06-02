package jira

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSprintsServiceCreate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/sprint", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		fmt.Fprint(w, `{"id": 5259,"self": "https://jira.mycompany.com/rest/agile/1.0/sprint/5259","state": "open","name": "Sprint 001","originBoardId": 2881,"goal": "My goal"}`)
	})

	newSprint := &NewSprint{
		Name:    "Sprint 001",
		BoardID: 2881,
	}

	sprint, _, err := client.Sprints.Create(context.Background(), newSprint)
	assert.Nil(t, err)

	assert.NotNil(t, sprint)
	assert.Equal(t, "Sprint 001", sprint.Name)
	assert.Equal(t, 2881, sprint.BoardID)
	assert.Equal(t, "My goal", sprint.Goal)
}

func TestSprintsServiceGet(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/sprint/5259", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		fmt.Fprint(w, `{"id": 5259,"self": "https://jira.mycompany.com/rest/agile/1.0/sprint/5259","state": "closed","name": "Sprint 001","startDate": "2018-09-18T17:30:00.000Z","endDate": "2018-09-19T01:30:00.000Z","completeDate": "2018-09-19T03:00:00.000Z","originBoardId": 2881,"goal": "My goal"}`)
	})

	sprint, _, err := client.Sprints.Get(context.Background(), 5259)
	assert.Nil(t, err)

	start := time.Date(2018, 9, 18, 17, 30, 0, 0, time.UTC)
	end := time.Date(2018, 9, 19, 1, 30, 0, 0, time.UTC)
	complete := time.Date(2018, 9, 19, 3, 0, 0, 0, time.UTC)

	want := &Sprint{
		ID:       5259,
		SelfLink: "https://jira.mycompany.com/rest/agile/1.0/sprint/5259",
		State:    "closed",
		Name:     "Sprint 001",
		Start:    &start,
		End:      &end,
		Complete: &complete,
		BoardID:  2881,
		Goal:     "My goal",
	}
	assert.True(t, reflect.DeepEqual(sprint, want))
}

func TestSprintsServiceUpdate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/sprint/11392", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method)
		fmt.Fprint(w, `{"id": 5259,"self": "https://jira.mycompany.com/rest/agile/1.0/sprint/5259","state": "open","name": "Sprint 001 XXX","originBoardId": 2881,"goal": "I do not know"}`)
	})

	newSprint := &Sprint{
		Name:    "Sprint 001 XXX",
		Goal:    "I do not know",
		BoardID: 2881,
	}

	sprint, _, err := client.Sprints.Update(context.Background(), 11392, newSprint)
	assert.Nil(t, err)

	assert.NotNil(t, sprint)
	assert.Equal(t, "Sprint 001 XXX", sprint.Name)
	assert.Equal(t, 2881, sprint.BoardID)
	assert.Equal(t, "I do not know", sprint.Goal)
}

func TestSprintsServicePartiallyUpdate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/sprint/11392", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		fmt.Fprint(w, `{"id": 5259,"self": "https://jira.mycompany.com/rest/agile/1.0/sprint/5259","state": "open","name": "Sprint 001 ---","originBoardId": 2881,"goal": "I do not know"}`)
	})

	newSprint := &Sprint{
		Name: "Sprint 001 ---",
	}

	sprint, _, err := client.Sprints.PartiallyUpdate(context.Background(), 11392, newSprint)
	assert.Nil(t, err)

	assert.NotNil(t, sprint)
	assert.Equal(t, "Sprint 001 ---", sprint.Name)
	assert.Equal(t, 2881, sprint.BoardID)
	assert.Equal(t, "I do not know", sprint.Goal)
}

func TestSprintsServiceMoveIssuesTo(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/sprint/5/issue", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusNoContent)
	})

	issues := &IssueKeys{
		Issues: []string{"MCP-1", "MCP-2"},
	}

	ok, _, err := client.Sprints.MoveIssuesTo(context.Background(), 5, issues)
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestSprintsServiceListIssuesForSprint(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/sprint/111/issue", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = fmt.Fprint(w, issuesAsJSON)
	})

	backlog, _, err := client.Sprints.ListIssues(context.Background(), 111, nil)
	assert.Nil(t, err)
	assert.Len(t, backlog, 1)
}
