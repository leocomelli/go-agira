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
