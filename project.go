package jira

// ProjectWrap represents the data returned by the API,
// in addition to the board information, paging data is returned
type ProjectWrap struct {
	Pagination
	Values []*Project `json:"values,omitempty"`
}

// ProjectCategory represents the project category of Jira Issue
type ProjectCategory struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	SelfLink    string `json:"self,omitempty"`
	Description string `json:"description,omitempty"`
}

// Project represents a Jira Project
type Project struct {
	ID         string            `json:"id,omitempty"`
	Key        string            `json:"key,omitempty"`
	Name       string            `json:"name,omitempty"`
	SelfLink   string            `json:"self,omitempty"`
	AvatarURLs map[string]string `json:"avatarUrls,omitempty"`
	Category   ProjectCategory   `json:"projectCategory,omitempty"`
	Simplified string            `json:"simplified,omitempty"`
	Style      string            `json:"style,omitempty"`
}

// ProjectsOptions contains all options to get a project from a board
type ProjectsOptions struct {
	//The starting index of the returned sprints. Base index: 0. See the 'Pagination' section at the top of this page for more details.
	StartAt int `query:"startAt"`
	//The maximum number of sprints to return per page. Default: 50. See the 'Pagination' section at the top of this page for more details.
	MaxResults int `query:"maxResults"`
}
