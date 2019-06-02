package jira

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// SprintsService handles communication with the sprint related
// methods of the Jira Agile API
//
// Jira Agile API docs: https://docs.atlassian.com/jira-software/REST/7.3.1/#agile/1.0/sprint
type SprintsService service

// SprintWrap represents the data returned by the API,
// in addition to the board information, paging data is returned
type SprintWrap struct {
	Pagination
	Values []*Sprint `json:"values,omitempty"`
}

// Sprint represents a Jira Agile Sprint
type Sprint struct {
	ID       int        `json:"id,omitempty"`
	Name     string     `json:"name,omitempty"`
	State    string     `json:"state,omitempty"`
	SelfLink string     `json:"self,omitempty"`
	Start    *time.Time `json:"startDate,omitempty"`
	End      *time.Time `json:"endDate,omitempty"`
	Complete *time.Time `json:"completeDate,omitempty"`
	BoardID  int        `json:"originBoardId,omitempty"`
	Goal     string     `json:"goal,omitempty"`
}

// NewSprint contains all options to create a sprint
type NewSprint struct {
	//Required
	Name    string `json:"name,omitempty"`
	BoardID int    `json:"originBoardId,omitempty"`
	//Optional
	Start *time.Time `json:"startDate,omitempty"`
	End   *time.Time `json:"endDate,omitempty"`
}

// SprintsOptions contains all options to list all sprints from a board
type SprintsOptions struct {
	//The starting index of the returned sprints. Base index: 0. See the 'Pagination' section at the top of this page for more details.
	StartAt int `query:"startAt"`
	//The maximum number of sprints to return per page. Default: 50. See the 'Pagination' section at the top of this page for more details.
	MaxResults int `query:"maxResults"`
	//Filters results to sprints in specified states. Valid values: future, active, closed. You can define multiple states separated by commas, e.g. state=active,closed
	State string `query:"state"`
}

// Create creates a future sprint. Sprint name and origin board id are required. Start and end date are optional.
// Note, the sprint name is trimmed.
//
// POST /rest/agile/1.0/sprint
func (s *SprintsService) Create(ctx context.Context, newSprint *NewSprint) (*Sprint, *Response, error) {

	req, err := s.client.NewRequest("POST", "sprint", newSprint)
	if err != nil {
		return nil, nil, err
	}

	var sprint = &Sprint{}
	resp, err := s.client.Do(ctx, req, sprint)
	if err != nil {
		return nil, resp, err
	}

	return sprint, resp, nil
}

// Get returns the sprint for a given sprint Id. The sprint will only be returned if the user can view
// the board that the sprint was created on, or view at least one of the issues in the sprint.
//
// GET /rest/agile/1.0/sprint/{sprintId}
func (s *SprintsService) Get(ctx context.Context, sprintID int) (*Sprint, *Response, error) {

	req, err := s.client.NewRequest("GET", fmt.Sprintf("sprint/%d", sprintID), nil)
	if err != nil {
		return nil, nil, err
	}

	var sprint = &Sprint{}
	resp, err := s.client.Do(ctx, req, sprint)
	if err != nil {
		return nil, resp, err
	}

	return sprint, resp, nil
}

// Update performs a full update of a sprint. A full update means that the result will be
// exactly the same as the request body. Any fields not present in the request JSON will
// be set to null.
//
// Notes:
// - Sprints that are in a closed state cannot be updated.
// - A sprint can be started by updating the state to 'active'. This requires the sprint to
// be in the 'future' state and have a startDate and endDate set.
// - A sprint can be completed by updating the state to 'closed'. This action requires the
// sprint to be in the 'active' state. This sets the completeDate to the time of the
// request.
// - Other changes to state are not allowed.
// - The completeDate field cannot be updated manually.
//
// PUT /rest/agile/1.0/sprint/{sprintId}
func (s *SprintsService) Update(ctx context.Context, sprintID int, sprintInfo *Sprint) (*Sprint, *Response, error) {

	req, err := s.client.NewRequest("PUT", fmt.Sprintf("sprint/%d", sprintID), sprintInfo)
	if err != nil {
		return nil, nil, err
	}

	var sprint = &Sprint{}
	resp, err := s.client.Do(ctx, req, sprint)
	if err != nil {
		return nil, resp, err
	}

	return sprint, resp, nil
}

// PartiallyUpdate performs a partial update of a sprint. A partial update means that fields not
// present in the request JSON will not be updated.
//
// Notes:
// - Sprints that are in a closed state cannot be updated.
// - A sprint can be started by updating the state to 'active'. This requires the sprint to be in the
// 'future' state and have a startDate and endDate set.
// - A sprint can be completed by updating the state to 'closed'. This action requires the sprint to
// be in the 'active' state. This sets the completeDate to the time of the request.
// - Other changes to state are not allowed.
// - The completeDate field cannot be updated manually.
//
// POST /rest/agile/1.0/sprint/{sprintId}
func (s *SprintsService) PartiallyUpdate(ctx context.Context, sprintID int, sprintInfo *Sprint) (*Sprint, *Response, error) {

	req, err := s.client.NewRequest("POST", fmt.Sprintf("sprint/%d", sprintID), sprintInfo)
	if err != nil {
		return nil, nil, err
	}

	var sprint = &Sprint{}
	resp, err := s.client.Do(ctx, req, sprint)
	if err != nil {
		return nil, resp, err
	}

	return sprint, resp, nil
}

// MoveIssuesTo Moves issues to a sprint, for a given sprint Id. Issues can only be moved to open or
// active sprints. The maximum number of issues that can be moved in one operation is 50.
//
// POST /rest/agile/1.0/sprint/{sprintId}/issue
func (s *SprintsService) MoveIssuesTo(ctx context.Context, sprintID int, issueKeys *IssueKeys) (bool, *Response, error) {

	req, err := s.client.NewRequest("POST", fmt.Sprintf("sprint/%d/issue", sprintID), issueKeys)
	if err != nil {
		return false, nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return false, resp, err
	}

	if resp.StatusCode == http.StatusNoContent {
		return true, resp, nil
	}

	return false, resp, nil
}
