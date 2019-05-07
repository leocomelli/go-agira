package jira

import (
	"context"
	"fmt"
	"time"
)

// SprintWrap represents the data returned by the API,
// in addition to the board information, paging data is returned
type SprintWrap struct {
	Pagination
	Values []*Sprint `json:"values,omitempty"`
}

// Sprint represents a Jira Agile Sprint
type Sprint struct {
	ID       int       `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	State    string    `json:"state,omitempty"`
	SelfLink string    `json:"self,omitempty"`
	Start    time.Time `json:"startDate,omitempty"`
	End      time.Time `json:"endDate,omitempty"`
	Complete time.Time `json:"completeDate,omitempty"`
	BoardID  int       `json:"originBoardId,omitempty"`
	Goal     string    `json:"goal,omitempty"`
}

func (s Sprint) String() string {
	return s.Name
}

// ListSprintsOptions contains all options to list all sprints from a board
type ListSprintsOptions struct {
	//The starting index of the returned sprints. Base index: 0. See the 'Pagination' section at the top of this page for more details.
	StartAt int `query:"startAt"`
	//The maximum number of sprints to return per page. Default: 50. See the 'Pagination' section at the top of this page for more details.
	MaxResults int `query:"maxResults"`
	//Filters results to sprints in specified states. Valid values: future, active, closed. You can define multiple states separated by commas, e.g. state=active,closed
	State string `query:"state"`
}

// ListSprints returns all sprints from a board, for a given board ID.
// This only includes sprints that the user has permission to view.
//
// GET /rest/agile/1.0/board/{boardId}/sprint
func (b *BoardsService) ListSprints(ctx context.Context, id int, opts *ListSprintsOptions) ([]*Sprint, *Response, error) {

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
