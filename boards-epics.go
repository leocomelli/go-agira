package jira

import (
	"context"
	"fmt"
)

// ListEpics returns all epics from the board, for the given board ID.
// This only includes epics that the user has permission to view.
// Note, if the user does not have permission to view the board,
// no epics will be returned at all.
//
// GET /rest/agile/1.0/board/{boardId}/epic
func (b *BoardsService) ListEpics(ctx context.Context, boardID int, opts *EpicsOptions) ([]*Epic, *Response, error) {

	q := QueryParameters(opts)

	req, err := b.client.NewRequest("GET", fmt.Sprintf("board/%d/epic%s", boardID, q), nil)
	if err != nil {
		return nil, nil, err
	}

	var wrap = &EpicWrap{}
	resp, err := b.client.Do(ctx, req, wrap)
	if err != nil {
		return nil, resp, err
	}

	resp.MaxResults = wrap.MaxResults
	resp.StartAt = wrap.StartAt
	resp.IsLast = wrap.IsLast

	return wrap.Values, resp, nil
}

// ListIssuesForEpic returns all issues that belong to an epic on the board,
// for the given epic Id and the board Id.
// This only includes issues that the user has permission to view. Issues
// returned from this resource include Agile fields, like sprint,
// closedSprints, flagged, and epic. By default, the returned issues are
// ordered by rank.
//
// GET /rest/agile/1.0/board/{boardId}/epic/{epicId}/issue
func (b *BoardsService) ListIssuesForEpic(ctx context.Context, id int, epicID int, opts *IssuesOptions) ([]*Issue, *Response, error) {

	q := QueryParameters(opts)

	req, err := b.client.NewRequest("GET", fmt.Sprintf("board/%d/epic/%d/issue%s", id, epicID, q), nil)
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

// ListIssuesWithoutEpic returns all issues that do not belong to any epic on a board,
// for a given board Id.
// This only includes issues that the user has permission to view. Issues returned
// from this resource include Agile fields, like sprint, closedSprints, flagged, and
// epic. By default, the returned issues are ordered by rank.
//
// GET /rest/agile/1.0/board/{boardId}/epic/none/issue
func (b *BoardsService) ListIssuesWithoutEpic(ctx context.Context, id int, opts *IssuesOptions) ([]*Issue, *Response, error) {

	q := QueryParameters(opts)

	req, err := b.client.NewRequest("GET", fmt.Sprintf("board/%d/epic/none/issue%s", id, q), nil)
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
