package jira

import (
	"context"
	"fmt"
)

// ListIssuesOptions contains all options to list backlog from a board
type ListIssuesOptions struct {
	//The starting index of the returned sprints. Base index: 0. See the 'Pagination' section at the top of this page for more details.
	StartAt int `query:"startAt"`
	//The maximum number of sprints to return per page. Default: 50. See the 'Pagination' section at the top of this page for more details.
	MaxResults int `query:"maxResults"`
	//Filters results using a JQL query. If you define an order in your JQL query, it will override the default order of the returned issues.
	JQL string `query:"jql"`
	//Specifies whether to validate the JQL query or not. Default: true.
	ValidateQuery bool `query:"validateQuery"`
	//The list of fields to return for each issue. By default, all navigable and Agile fields are returned.
	Fields string `query:"fields"`
	//This parameter is currently not used.
	Expand string `query:"expand"`
}

// ListBacklogIssues returns all issues from the board's backlog, for the given board Id.
// This only includes issues that the user has permission to view. The backlog contains
// incomplete issues that are not assigned to any future or active sprint. Note, if the
// user does not have permission to view the board, no issues will be returned at all.
// Issues returned from this resource include Agile fields, like sprint, closedSprints,
// flagged, and epic. By default, the returned issues are ordered by rank.
//
// GET /rest/agile/1.0/board/{boardId}/backlog
func (b *BoardsService) ListBacklogIssues(ctx context.Context, id int, opts *ListIssuesOptions) ([]*Issue, *Response, error) {

	q := QueryParameters(opts)

	req, err := b.client.NewRequest("GET", fmt.Sprintf("board/%d/backlog%s", id, q), nil)
	if err != nil {
		return nil, nil, err
	}

	var wrap = &IssueWrap{}
	resp, err := b.client.Do(ctx, req, wrap)
	if err != nil {
		return nil, resp, err
	}

	resp.MaxResults = wrap.MaxResults
	resp.StartAt = wrap.StartAt
	resp.IsLast = wrap.IsLast

	return wrap.Values, resp, nil
}

// ListIssues returns all issues from a board, for a given board Id.
// This only includes issues that the user has permission to view. Note,
// if the user does not have permission to view the board, no issues will
// be returned at all. Issues returned from this resource include Agile
// fields, like sprint, closedSprints, flagged, and epic.
// By default, the returned issues are ordered by rank.
//
// GET /rest/agile/1.0/board/{boardId}/issue
func (b *BoardsService) ListIssues(ctx context.Context, id int, opts *ListIssuesOptions) ([]*Issue, *Response, error) {

	q := QueryParameters(opts)

	req, err := b.client.NewRequest("GET", fmt.Sprintf("board/%d/issue%s", id, q), nil)
	if err != nil {
		return nil, nil, err
	}

	var wrap = &IssueWrap{}
	resp, err := b.client.Do(ctx, req, wrap)
	if err != nil {
		return nil, resp, err
	}

	resp.MaxResults = wrap.MaxResults
	resp.StartAt = wrap.StartAt
	resp.IsLast = wrap.IsLast

	return wrap.Values, resp, nil
}
