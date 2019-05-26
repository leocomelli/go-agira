package jira

import (
	"context"
	"fmt"
)

// ListVersions returns all versions from a board, for a given board Id.
// This only includes versions that the user has permission to view.
// Note, if the user does not have permission to view the board, no
// versions will be returned at all. Returned versions are ordered by
// the name of the project from which they belong and then by sequence
// defined by user.
//
// GET /rest/agile/1.0/board/{boardId}/version
func (b *BoardsService) ListVersions(ctx context.Context, id int, opts *VersionsOptions) ([]*Version, *Response, error) {

	q := QueryParameters(opts)

	req, err := b.client.NewRequest("GET", fmt.Sprintf("board/%d/version%s", id, q), nil)
	if err != nil {
		return nil, nil, err
	}

	var wrap = &VersionWrap{}
	resp, err := b.client.Do(ctx, req, wrap)
	if err != nil {
		return nil, resp, err
	}

	resp.MaxResults = wrap.MaxResults
	resp.StartAt = wrap.StartAt
	resp.IsLast = wrap.IsLast

	return wrap.Values, resp, nil
}
