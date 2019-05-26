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

	epic, _, err := client.Epics.GetEpic(context.Background(), "5")
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
