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

	want := &Sprint{
		ID:       5259,
		SelfLink: "https://jira.mycompany.com/rest/agile/1.0/sprint/5259",
		State:    "closed",
		Name:     "Sprint 001",
		Start:    time.Date(2018, 9, 18, 17, 30, 0, 0, time.UTC),
		End:      time.Date(2018, 9, 19, 1, 30, 0, 0, time.UTC),
		Complete: time.Date(2018, 9, 19, 3, 0, 0, 0, time.UTC),
		BoardID:  2881,
		Goal:     "My goal",
	}
	assert.True(t, reflect.DeepEqual(sprint, want))
}
