package jira

import (
	"context"
)

// BoardsService handles communication with the issue related
// methods of the Jira Agile API
//
// Jira Agile API docs: https://developer.atlassian.com/cloud/jira/software/rest/#api-rest-agile-1-0-board-get
type BoardsService service

// BoardWrap represents the data returned by the API,
// in addition to the board information, paging data is returned
type BoardWrap struct {
	Pagination
	Values []*Board `json:"values,omitempty"`
}

// Board represents a Jira Agile Board
type Board struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Type     string `json:"type,omitempty"`
	SelfLink string `json:"self,omitempty"`
}

func (b Board) String() string {
	return b.Name
}

// ListBoardsOptions contains all options to list boards
type ListBoardsOptions struct {
	//The starting index of the returned boards. Base index: 0. See the 'Pagination' section at the top of this page for more details.
	StartAt int `query:"startAt"`
	//The maximum number of boards to return per page. Default: 50. See the 'Pagination' section at the top of this page for more details.
	MaxResults int `query:"maxResults"`
	//Filters results to boards of the specified types. Valid values: scrum, kanban, simple.
	Type string `query:"type"`
	//Filters results to boards that match or partially match the specified name.
	Name string `query:"name"`
	//Filters results to boards that are relevant to a project. Relevance means that the jql filter defined in board contains a reference to a project.
	ProjectKeyOrID string `query:"projectKeyOrId"`
	//Appends private boards to the end of the list. The name and type fields are excluded for security reasons.
	IncludePrivate bool `query:"includePrivate"`
	//If set to true, negate filters used for querying by location. By default false.
	NegateLocationFiltering bool `query:"negateLocationFiltering"`
	//Ordering of the results by a given field. If not provided, values will not be sorted. Valid values: name.
	OrderBy string `query:"orderBy"`
	//List of fields to expand for each board. Valid values: admins, permissions.
	Expand string `query:"expand"`
	//Filters results to boards that are relevant to a filter. Not supported for next-gen boards.
	FilterID int `query:"filterId"`

	AccountIDLocation string `query:"accountIdLocation"`
	UserKeyLocation   string `query:"userkeyLocation"`
	UsernameLocation  string `query:"usernameLocation"`
	ProjectLocation   string `query:"projectLocation"`
}

// ListBoards returns all boards
// This only includes boards that the user has permission to view
//
// GET /rest/agile/1.0/board
func (b *BoardsService) ListBoards(ctx context.Context, opts *ListBoardsOptions) ([]*Board, *Response, error) {

	q := QueryParameters(opts)

	req, err := b.client.NewRequest("GET", "board"+q, nil)
	if err != nil {
		return nil, nil, err
	}

	var wrap = &BoardWrap{}
	resp, err := b.client.Do(ctx, req, wrap)
	if err != nil {
		return nil, resp, err
	}

	resp.MaxResults = wrap.MaxResults
	resp.StartAt = wrap.StartAt
	resp.IsLast = wrap.IsLast

	return wrap.Values, resp, nil
}
