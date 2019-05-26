package jira

// VersionWrap represents the data returned by the API,
// in addition to the board information, paging data is returned
type VersionWrap struct {
	Pagination
	Values []*Version `json:"values,omitempty"`
}

// Version represents the version of Jira Issue
type Version struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	SelfLink    string `json:"self,omitempty"`
	Description string `json:"description,omitempty"`
	Archived    bool   `json:"archived,omitempty"`
	Released    bool   `json:"released,omitempty"`
	ProjectID   int    `json:"projectId,omitempty"`
}

// VersionsOptions contains all options to list all versions from the board
type VersionsOptions struct {
	//The starting index of the returned epics. Base index: 0. See the 'Pagination' section at the top of this page for more details.
	StartAt int `query:"startAt"`
	//The maximum number of epics to return per page. Default: 50. See the 'Pagination' section at the top of this page for more details.
	MaxResults int `query:"maxResults"`
	//Filters results to versions that are either released or unreleased. Valid values: true, false.
	Released string `query:"released"`
}
