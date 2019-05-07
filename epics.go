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

func (e Epic) String() string {
	return e.Name
}
