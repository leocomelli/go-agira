package jira

import "time"

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

// SprintsOptions contains all options to list all sprints from a board
type SprintsOptions struct {
	//The starting index of the returned sprints. Base index: 0. See the 'Pagination' section at the top of this page for more details.
	StartAt int `query:"startAt"`
	//The maximum number of sprints to return per page. Default: 50. See the 'Pagination' section at the top of this page for more details.
	MaxResults int `query:"maxResults"`
	//Filters results to sprints in specified states. Valid values: future, active, closed. You can define multiple states separated by commas, e.g. state=active,closed
	State string `query:"state"`
}
