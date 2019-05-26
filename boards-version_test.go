package jira

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoardsServiceListVersions(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/board/5259/version", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = fmt.Fprint(w, `{"maxResults": 50,"startAt": 0,"isLast": true,
"values": [{"self": "https://jira.mycompany.com/rest/api/2/version/28059","id": 28059,"projectId": 17526,"name": "1.0.0","description": "","archived": false,"released": false}]}`)
	})

	sprints, resp, err := client.Boards.ListVersions(context.Background(), 5259, nil)
	assert.Nil(t, err)
	assert.Len(t, sprints, 1)

	want := []*Version{
		{
			ID:          28059,
			SelfLink:    "https://jira.mycompany.com/rest/api/2/version/28059",
			Name:        "1.0.0",
			Description: "",
			ProjectID:   17526,
			Archived:    false,
			Released:    false,
		},
	}
	assert.True(t, reflect.DeepEqual(sprints, want))
	assert.Equal(t, 50, resp.MaxResults)
	assert.Equal(t, 0, resp.StartAt)
	assert.True(t, resp.IsLast)
}
