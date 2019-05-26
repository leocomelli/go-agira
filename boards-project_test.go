package jira

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoardsServiceListProject(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/board/5259/project", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = fmt.Fprint(w, `{"maxResults": 50,"startAt": 0,"total": 1,"isLast": true,
        "values": [{"self": "https://jira.company.com/rest/api/2/project/17526","id": "17526","key": "CBD","name": "Digital","avatarUrls": {"48x48": "https://jira.company.com/secure/projectavatar?pid=17526&avatarId=20500","24x24": "https://jira.company.com/secure/projectavatar?size=small&pid=17526&avatarId=20500","16x16": "https://jira.company.com/secure/projectavatar?size=xsmall&pid=17526&avatarId=20500","32x32": "https://jira.company.com/secure/projectavatar?size=medium&pid=17526&avatarId=20500"}}]}`)
	})

	sprints, resp, err := client.Boards.ListProjects(context.Background(), 5259, nil)
	assert.Nil(t, err)
	assert.Len(t, sprints, 1)

	want := []*Project{
		{
			ID:       "17526",
			SelfLink: "https://jira.company.com/rest/api/2/project/17526",
			Name:     "Digital",
			Key:      "CBD",
			AvatarURLs: map[string]string{
				"48x48": "https://jira.company.com/secure/projectavatar?pid=17526&avatarId=20500",
				"24x24": "https://jira.company.com/secure/projectavatar?size=small&pid=17526&avatarId=20500",
				"16x16": "https://jira.company.com/secure/projectavatar?size=xsmall&pid=17526&avatarId=20500",
				"32x32": "https://jira.company.com/secure/projectavatar?size=medium&pid=17526&avatarId=20500",
			},
		},
	}
	assert.True(t, reflect.DeepEqual(sprints, want))
	assert.Equal(t, 50, resp.MaxResults)
	assert.Equal(t, 0, resp.StartAt)
	assert.True(t, resp.IsLast)
}
