package jira

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoardsServiceList(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/board", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		fmt.Fprint(w, `{"maxResults": 50,"startAt": 0,"isLast": false,
		"values": [{"id": 42,"self": "https://jira.com/rest/agile/1.0/board/42","name": "MTD board","type": "scrum"}]}`)
	})

	boards, resp, err := client.Boards.ListBoards(context.Background(), nil)
	assert.Nil(t, err)
	assert.Len(t, boards, 1)

	want := []*Board{
		{
			ID:       42,
			SelfLink: "https://jira.com/rest/agile/1.0/board/42",
			Name:     "MTD board",
			Type:     "scrum",
		},
	}
	assert.True(t, reflect.DeepEqual(boards, want))
	assert.Equal(t, 50, resp.MaxResults)
	assert.Equal(t, 0, resp.StartAt)
	assert.False(t, resp.IsLast)
}
