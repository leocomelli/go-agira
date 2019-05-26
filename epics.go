package jira

// EpicWrap represents the data returned by the API,
// in addition to the board information, paging data is returned
type EpicWrap struct {
	Pagination
	Values []*Epic `json:"values,omitempty"`
}

// Epic represents a Jira Agile Epic
type Epic struct {
	ID       int               `json:"id,omitempty"`
	Key      string            `json:"key,omitempty"`
	Name     string            `json:"name,omitempty"`
	Summary  string            `json:"summary,omitempty"`
	SelfLink string            `json:"self,omitempty"`
	Done     bool              `json:"done,omitempty"`
	Color    map[string]string `json:"color,omitempty"`
}

// EpicsOptions contains all options to list all epics from the board
type EpicsOptions struct {
	//The starting index of the returned epics. Base index: 0. See the 'Pagination' section at the top of this page for more details.
	StartAt int `query:"startAt"`
	//The maximum number of epics to return per page. Default: 50. See the 'Pagination' section at the top of this page for more details.
	MaxResults int `query:"maxResults"`
	//Filters results to epics that are either done or not done. Valid values: true, false.
	Done bool `query:"done"`
}
