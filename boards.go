package jira

import (
	"context"
	"fmt"
)

// BoardsService handles communication with the board related
// methods of the Jira Agile API
//
// Jira Agile API docs: https://docs.atlassian.com/jira-software/REST/7.3.1/#agile/1.0/board
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

// NewBoard contains all options to create a board
type NewBoard struct {
	//Must be less than 255 characters.
	Name string `json:"name,omitempty"`
	//Valid values: scrum, kanban
	Type string `json:"type,omitempty"`
	//Id of a filter that the user has permissions to view. Note, if the user does not have the
	// 'Create shared objects' permission and tries to create a shared board, a private board will
	// be created instead (remember that board sharing depends on the filter sharing).
	FilterID int `json:"filterId,omitempty"`
}

// BoardsOptions contains all options to list boards
type BoardsOptions struct {
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

// CreateBoard creates a new board. Board name, type and filter Id is required.
//
// POST /rest/agile/1.0/board
func (b *BoardsService) CreateBoard(ctx context.Context, newBoard *NewBoard) (*Board, *Response, error) {

	req, err := b.client.NewRequest("POST", "board", newBoard)
	if err != nil {
		return nil, nil, err
	}

	var board = &Board{}
	resp, err := b.client.Do(ctx, req, board)
	if err != nil {
		return nil, resp, err
	}

	return board, resp, nil
}

// DeleteBoard deletes the board.
//
// DELETE /rest/agile/1.0/board/{boardId}
func (b *BoardsService) DeleteBoard(ctx context.Context, id int) (*Response, error) {

	req, err := b.client.NewRequest("DELETE", fmt.Sprintf("board/%d", id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := b.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ListBoards returns all boards
// This only includes boards that the user has permission to view
//
// GET /rest/agile/1.0/board
func (b *BoardsService) ListBoards(ctx context.Context, opts *BoardsOptions) ([]*Board, *Response, error) {

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

// GetBoard returns the board for the given board Id.
// This board will only be returned if the user has permission to view it.
//
// GET /rest/agile/1.0/board/{boardId}
func (b *BoardsService) GetBoard(ctx context.Context, boardID int) (*Board, *Response, error) {

	req, err := b.client.NewRequest("GET", fmt.Sprintf("board/%d", boardID), nil)
	if err != nil {
		return nil, nil, err
	}

	var board = &Board{}
	resp, err := b.client.Do(ctx, req, board)
	if err != nil {
		return nil, resp, err
	}

	return board, resp, nil
}

// ListBacklogIssues returns all issues from the board's backlog, for the given board Id.
// This only includes issues that the user has permission to view. The backlog contains
// incomplete issues that are not assigned to any future or active sprint. Note, if the
// user does not have permission to view the board, no issues will be returned at all.
// Issues returned from this resource include Agile fields, like sprint, closedSprints,
// flagged, and epic. By default, the returned issues are ordered by rank.
//
// GET /rest/agile/1.0/board/{boardId}/backlog
func (b *BoardsService) ListBacklogIssues(ctx context.Context, id int, opts *IssuesOptions) ([]*Issue, *Response, error) {

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
func (b *BoardsService) ListIssues(ctx context.Context, id int, opts *IssuesOptions) ([]*Issue, *Response, error) {

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
