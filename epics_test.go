package jira

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEpicsServiceGet(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/epic/5", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		fmt.Fprint(w, `{"id": 523967,"key": "MCP-9","self": "https://jira.mycompany.com/rest/agile/1.0/epic/523967","name": "Epic 1","summary": "Epic 1","color": {"key": "color_9"},"done": false}`)
	})

	epic, _, err := client.Epics.Get(context.Background(), "5")
	assert.Nil(t, err)

	want := &Epic{
		ID:       523967,
		Key:      "MCP-9",
		SelfLink: "https://jira.mycompany.com/rest/agile/1.0/epic/523967",
		Name:     "Epic 1",
		Summary:  "Epic 1",
		Color: map[string]string{
			"key": "color_9",
		},
		Done: false,
	}

	assert.True(t, reflect.DeepEqual(epic, want))
}

func TestEpicsServiceListIssues(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/epic/5259/issue", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = fmt.Fprint(w, issueAsJSON)
	})

	backlog, _, err := client.Epics.ListIssues(context.Background(), "5259", nil)
	assert.Nil(t, err)
	assert.Len(t, backlog, 1)
}

func TestEpicsServicePartiallyUpdate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/epic/5", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		fmt.Fprint(w, `{"id": 523967,"key": "MCP-9","self": "https://jira.mycompany.com/rest/agile/1.0/epic/523967","name": "Epic 1","summary": "Epic 1","color": {"key": "color_9"},"done": false}`)
	})

	epic, _, err := client.Epics.PartiallyUpdate(context.Background(), "5", &Epic{})
	assert.Nil(t, err)

	want := &Epic{
		ID:       523967,
		Key:      "MCP-9",
		SelfLink: "https://jira.mycompany.com/rest/agile/1.0/epic/523967",
		Name:     "Epic 1",
		Summary:  "Epic 1",
		Color: map[string]string{
			"key": "color_9",
		},
		Done: false,
	}

	assert.True(t, reflect.DeepEqual(epic, want))
}

func TestEpicsServiceMoveIssuesTo(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/epic/5/issue", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusNoContent)
	})

	issues := &IssueKeys{
		Issues: []string{"MCP-1", "MCP-2"},
	}

	ok, err := client.Epics.MoveIssuesTo(context.Background(), "5", issues)
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestEpicsServiceRemoveIssuesFrom(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/epic/none/issue", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusNoContent)
	})

	issues := &IssueKeys{
		Issues: []string{"MCP-1", "MCP-2"},
	}

	ok, err := client.Epics.RemoveIssuesFrom(context.Background(), issues)
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestEpicsServiceListIssuesWithoutEpic(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/epic/none/issue", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = fmt.Fprint(w, issueAsJSON)
	})

	backlog, _, err := client.Epics.ListIssuesWithoutEpic(context.Background(), nil)
	assert.Nil(t, err)
	assert.Len(t, backlog, 1)
}
