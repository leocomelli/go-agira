package jira

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBacklogServiceMoveIssuesTo(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/backlog/issue", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusNoContent)
	})

	issues := &IssueKeys{
		Issues: []string{"MCP-1", "MCP-2"},
	}

	ok, _, err := client.Backlog.MoveIssuesTo(context.Background(), issues)
	assert.Nil(t, err)
	assert.True(t, ok)
}
