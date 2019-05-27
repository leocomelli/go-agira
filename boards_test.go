package jira

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoardsServiceCreate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/board", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		fmt.Fprint(w, `{"id": 5597,"self": "https://jira.com/rest/agile/1.0/board/42","name": "MTD board","type": "scrum"}`)

	})

	b, _, err := client.Boards.Create(context.Background(), &NewBoard{})
	assert.Nil(t, err)
	assert.Equal(t, b.ID, 5597)
}

func TestBoardsServiceDelete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/board/9999", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method)
	})

	_, err := client.Boards.Delete(context.Background(), 9999)
	assert.Nil(t, err)
}

func TestBoardsServiceList(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/board", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		fmt.Fprint(w, `{"maxResults": 50,"startAt": 0,"isLast": false,
		"values": [{"id": 42,"self": "https://jira.com/rest/agile/1.0/board/42","name": "MTD board","type": "scrum"}]}`)
	})

	boards, resp, err := client.Boards.List(context.Background(), nil)
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

func TestBoardsServiceGet(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/board/5597", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		fmt.Fprint(w, `{"id": 5597,"self": "https://jira.com/rest/agile/1.0/board/42","name": "MTD board","type": "scrum"}`)
	})

	board, _, err := client.Boards.Get(context.Background(), 5597)
	assert.Nil(t, err)

	want := &Board{
		ID:       5597,
		SelfLink: "https://jira.com/rest/agile/1.0/board/42",
		Name:     "MTD board",
		Type:     "scrum",
	}

	assert.True(t, reflect.DeepEqual(board, want))
}

func TestBoardsServiceListBacklogIssues(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/board/5259/backlog", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = fmt.Fprint(w, issueAsJSON)
	})

	backlog, resp, err := client.Boards.ListBacklogIssues(context.Background(), 5259, nil)
	assert.Nil(t, err)
	assert.Len(t, backlog, 1)

	issue := backlog[0]
	assert.Equal(t, "776509", issue.ID)
	assert.Equal(t, "https://jira.mycompany.com/rest/agile/1.0/issue/776509", issue.SelfLink)
	assert.Equal(t, "MCP-840", issue.Key)
	assert.NotNil(t, issue.Fields.Project)
	assert.Len(t, issue.Fields.Worklogs.Worklogs, 4)
	assert.Len(t, issue.Fields.Comments.Comments, 2)
	assert.Equal(t, 50, resp.MaxResults)
	assert.Equal(t, 0, resp.StartAt)
	assert.False(t, resp.IsLast)
}

func TestBoardsServiceListIssues(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/board/5259/issue", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = fmt.Fprint(w, issueAsJSON)
	})

	backlog, _, err := client.Boards.ListIssues(context.Background(), 5259, nil)
	assert.Nil(t, err)
	assert.Len(t, backlog, 1)
}

func TestBoardsServiceGetConfiguration(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/board/5597/configuration", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		fmt.Fprint(w, `{
			"id": 5597,
			"name": "Board Configuration",
			"self": "http://www.example.com/jira/rest/agile/1.0/board/84/config",
			"filter": {
				"id": "1001",
				"self": "http://www.example.com/jira/filter/1001"
			},
			"columnConfig": {
				"columns": [
					{
						"name": "To Do",
						"statuses": [
							{
								"id": "1",
								"self": "http://www.example.com/jira/status/1"
							},
							{
								"id": "4",
								"self": "http://www.example.com/jira/status/4"
							}
						]
					},
					{
						"name": "In progress",
						"statuses": [
							{
								"id": "3",
								"self": "http://www.example.com/jira/status/3"
							}
						],
						"min": 2,
						"max": 4
					},
					{
						"name": "Done",
						"statuses": [
							{
								"id": "5",
								"self": "http://www.example.com/jira/status/5"
							}
						]
					}
				],
				"constraintType": "issueCount"
			},
			"estimation": {
				"type": "field",
				"field": {
					"fieldId": "customfield_10002",
					"displayName": "Story Points"
				}
			},
			"ranking": {
				"rankCustomFieldId": 10020
			}
		}`)
	})

	configuration, _, err := client.Boards.GetConfiguration(context.Background(), 5597)
	assert.Nil(t, err)

	want := &Configuration{
		ID:       5597,
		SelfLink: "http://www.example.com/jira/rest/agile/1.0/board/84/config",
		Name:     "Board Configuration",
		Filter: ConfigurationFilter{
			ID:       "1001",
			SelfLink: "http://www.example.com/jira/filter/1001",
		},
		Estimation: ConfigurationEstimation{
			Type: "field",
			Field: ConfigurationEstimationField{
				ID:   "customfield_10002",
				Name: "Story Points",
			},
		},
	}

	assert.Equal(t, want.ID, configuration.ID)
	assert.Equal(t, want.Name, configuration.Name)
	assert.Equal(t, want.Estimation.Field.ID, configuration.Estimation.Field.ID)

}
