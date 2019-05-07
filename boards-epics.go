package jira

import (
	"context"
	"fmt"
)

// ListEpicsOptions contains all options to list all epics from the board
type ListEpicsOptions struct {
	//The starting index of the returned epics. Base index: 0. See the 'Pagination' section at the top of this page for more details.
	StartAt int
	//The maximum number of epics to return per page. Default: 50. See the 'Pagination' section at the top of this page for more details.
	MaxResults int
	//Filters results to epics that are either done or not done. Valid values: true, false.
	Done string
}

// ListEpics returns all epics from the board, for the given board ID.
// This only includes epics that the user has permission to view.
// Note, if the user does not have permission to view the board,
// no epics will be returned at all.
//
// GET /rest/agile/1.0/board/{boardId}/epic
func (b *BoardsService) ListEpics(ctx context.Context, boardID int, opts *ListEpicsOptions) ([]*Epic, *Response, error) {

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
