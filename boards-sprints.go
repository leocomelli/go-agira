package jira

import (
	"context"
	"fmt"
)

// ListSprints returns all sprints from a board, for a given board ID.
// This only includes sprints that the user has permission to view.
//
// GET /rest/agile/1.0/board/{boardId}/sprint
func (b *BoardsService) ListSprints(ctx context.Context, id int, opts *SprintsOptions) ([]*Sprint, *Response, error) {

	q := QueryParameters(opts)

	req, err := b.client.NewRequest("GET", fmt.Sprintf("board/%d/sprint%s", id, q), nil)
	if err != nil {
		return nil, nil, err
	}

	var wrap = &SprintWrap{}
	resp, err := b.client.Do(ctx, req, wrap)
	if err != nil {
		return nil, resp, err
	}

	resp.MaxResults = wrap.MaxResults
	resp.StartAt = wrap.StartAt
	resp.IsLast = wrap.IsLast

	return wrap.Values, resp, nil
}

// ListIssuesForSprint get all issues you have access to that belong to the sprint
// from the board. Issue returned from this resource contains additional fields like:
// sprint, closedSprints, flagged and epic. Issues are returned ordered by rank.
// JQL order has higher priority than default rank.
//
// GET /rest/agile/1.0/board/{boardId}/sprint/{sprintId}/issue
func (b *BoardsService) ListIssuesForSprint(ctx context.Context, id int, sprintID int, opts *IssuesOptions) ([]*Issue, *Response, error) {
	q := QueryParameters(opts)

	req, err := b.client.NewRequest("GET", fmt.Sprintf("board/%d/sprint/%d/issue%s", id, sprintID, q), nil)
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
