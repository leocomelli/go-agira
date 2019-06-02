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

func TestBoardsServiceListSprints(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/board/5259/sprint", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		fmt.Fprint(w, `{"maxResults": 50,"startAt": 0,"isLast": false,
		"values": [{"id": 5259,"self": "https://jira.mycompany.com/rest/agile/1.0/sprint/5259","state": "closed","name": "Sprint 001","startDate": "2018-09-18T17:30:00.000Z","endDate": "2018-09-19T01:30:00.000Z","completeDate": "2018-09-19T03:00:00.000Z","originBoardId": 2881,"goal": "My goal"}]}`)
	})

	sprints, resp, err := client.Boards.ListSprints(context.Background(), 5259, nil)
	assert.Nil(t, err)
	assert.Len(t, sprints, 1)

	want := []*Sprint{
		{
			ID:       5259,
			SelfLink: "https://jira.mycompany.com/rest/agile/1.0/sprint/5259",
			State:    "closed",
			Name:     "Sprint 001",
			Start:    time.Date(2018, 9, 18, 17, 30, 0, 0, time.UTC),
			End:      time.Date(2018, 9, 19, 1, 30, 0, 0, time.UTC),
			Complete: time.Date(2018, 9, 19, 3, 0, 0, 0, time.UTC),
			BoardID:  2881,
			Goal:     "My goal",
		},
	}
	assert.True(t, reflect.DeepEqual(sprints, want))
	assert.Equal(t, 50, resp.MaxResults)
	assert.Equal(t, 0, resp.StartAt)
	assert.False(t, resp.IsLast)
}

func TestBoardsServiceListIssuesForSprint(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/board/5259/sprint/111/issue", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = fmt.Fprint(w, issuesAsJSON)
	})

	backlog, _, err := client.Boards.ListIssuesForSprint(context.Background(), 5259, 111, nil)
	assert.Nil(t, err)
	assert.Len(t, backlog, 1)
}
