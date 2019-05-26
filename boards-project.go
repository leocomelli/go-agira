package jira

import (
	"context"
	"fmt"
)

// ListProjects returns all projects that are associated with the board,
// for the given board Id. A project is associated with a board only if
// the board filter explicitly filters issues by the project and guaranties
// that all issues will come for one of those projects e.g. board's filter
// with "project in (PR-1, PR-1) OR reporter = admin" jql Projects are
// returned only if user can browse all projects that are associated with
// the board. Note, if the user does not have permission to view the board,
// no projects will be returned at all. Returned projects are ordered by the name.
//
// GET /rest/agile/1.0/board/{boardId}/project
func (b *BoardsService) ListProjects(ctx context.Context, id int, opts *ProjectsOptions) ([]*Project, *Response, error) {

	q := QueryParameters(opts)

	req, err := b.client.NewRequest("GET", fmt.Sprintf("board/%d/project%s", id, q), nil)
	if err != nil {
		return nil, nil, err
	}

	var wrap = &ProjectWrap{}
	resp, err := b.client.Do(ctx, req, wrap)
	if err != nil {
		return nil, resp, err
	}

	resp.MaxResults = wrap.MaxResults
	resp.StartAt = wrap.StartAt
	resp.IsLast = wrap.IsLast

	return wrap.Values, resp, nil
}
