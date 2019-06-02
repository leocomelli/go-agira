package jira

import (
	"context"
	"fmt"
	"net/http"
	"testing"

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
