package jira

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoardsServiceListEpics(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/board/5259/epic", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		fmt.Fprint(w, `{"maxResults": 50,"startAt": 0,"isLast": false,
		"values": [{"id": 523967,"key": "CBD-9","self": "https://jira.mycompany.com/rest/agile/1.0/epic/523967","name": "Order","summary": "Order","color": {"key": "color_9"},"done": false}]}`)
	})

	sprints, resp, err := client.Boards.ListEpics(context.Background(), 5259, nil)
	assert.Nil(t, err)
	assert.Len(t, sprints, 1)

	want := []*Epic{
		{
			ID:       523967,
			Key:      "CBD-9",
			SelfLink: "https://jira.mycompany.com/rest/agile/1.0/epic/523967",
			Name:     "Order",
			Summary:  "Order",
			Color:    map[string]string{"key": "color_9"},
			Done:     false,
		},
	}
	assert.True(t, reflect.DeepEqual(sprints, want))
	assert.Equal(t, 50, resp.MaxResults)
	assert.Equal(t, 0, resp.StartAt)
	assert.False(t, resp.IsLast)
}

func TestBoardsServiceListIssuesForEpic(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/board/5259/epic/1/issue", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = fmt.Fprint(w, issuesAsJSON)
	})

	backlog, _, err := client.Boards.ListIssuesForEpic(context.Background(), 5259, 1, nil)
	assert.Nil(t, err)
	assert.Len(t, backlog, 1)
}

func TestBoardsServiceListIssuesWithoutEpic(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/board/5259/epic/none/issue", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = fmt.Fprint(w, issuesAsJSON)
	})

	backlog, _, err := client.Boards.ListIssuesWithoutEpic(context.Background(), 5259, nil)
	assert.Nil(t, err)
	assert.Len(t, backlog, 1)
}
