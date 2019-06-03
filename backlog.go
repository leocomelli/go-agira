package jira

import (
	"context"
	"net/http"
)

// BacklogService handles communication with the backlog related
// methods of the Jira Agile API
//
// Jira Agile API docs: https://docs.atlassian.com/jira-software/REST/7.3.1/#agile/1.0/backlog
type BacklogService service

// MoveIssuesTo move issues to the backlog.
// This operation is equivalent to remove future and active sprints from a given set of issues.
// At most 50 issues may be moved at once.
//
// POST /rest/agile/1.0/backlog/issue
func (b *BacklogService) MoveIssuesTo(ctx context.Context, issueKeys *IssueKeys) (bool, *Response, error) {
	req, err := b.client.NewRequest("POST", "backlog/issue", issueKeys)
	if err != nil {
		return false, nil, err
	}

	resp, err := b.client.Do(ctx, req, nil)
	if err != nil {
		return false, resp, err
	}

	if resp.StatusCode == http.StatusNoContent {
		return true, resp, nil
	}

	return false, resp, nil
}
