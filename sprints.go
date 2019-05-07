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

func (s Sprint) String() string {
	return s.Name
}
