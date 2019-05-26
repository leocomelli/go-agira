package jira

import (
	"encoding/json"
	"strings"
	"time"
)

// DateTime represents a time in 2006-01-02T15:04:05.000-0700 format
type DateTime time.Time

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in 2006-01-02T15:04:05.000-0700 format.
func (d *DateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")

	if s == "null" {
		*d = DateTime(time.Time{})
		return nil
	}

	t, err := time.Parse("2006-01-02T15:04:05.000-0700", s)
	if err != nil {
		return err
	}
	*d = DateTime(t)
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
// The time is a quoted string in 2006-01-02T15:04:05.000-0700 format
func (d DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(d)
}

// IssueWrap represents the data returned by the API,
// in addition to the issues information, paging data is returned
type IssueWrap struct {
	Pagination
	Expand string   `json:"expand,omitempty"`
	Values []*Issue `json:"issues,omitempty"`
}

// Issue represents a Jira Issue
type Issue struct {
	ID       string      `json:"id,omitempty"`
	Key      string      `json:"key,omitempty"`
	SelfLink string      `json:"self,omitempty"`
	Expand   string      `json:"expand,omitempty"`
	Fields   *IssueField `json:"fields,omitempty"`
}

// IssueField represents the fields of Jira Issue
type IssueField struct {
	Flagged                       bool               `json:"flagged,omitempty"`
	Description                   string             `json:"description,omitempty"`
	Sprint                        *Sprint            `json:"sprint,omitempty"`
	ClosedSprints                 []*Sprint          `json:"closedSprints,omitempty"`
	Project                       *IssueProject      `json:"project,omitempty"`
	Resolution                    *IssueResolution   `json:"resolution,omitempty"`
	LastViewed                    DateTime           `json:"lastViewed,omitempty"`
	AggregateTimeOriginalEstimate int                `json:"aggregatetimeoriginalestimate,omitempty"`
	AggregateTimeEstimate         int                `json:"aggregatetimeestimate,omitempty"`
	Links                         []IssueLink        `json:"issuelinks,omitempty"`
	SubTasks                      []*Issue           `json:"subtasks,omitempty"`
	Type                          IssueType          `json:"issuetype,omitempty"`
	Environment                   string             `json:"environment,omitempty"`
	TimeEstimate                  int                `json:"timeestimate,omitempty"`
	AggregateTimeSpent            int                `json:"aggregatetimespent,omitempty"`
	WorkRatio                     int                `json:"workratio,omitempty"`
	Labels                        []string           `json:"labels,omitempty"`
	Reporter                      *IssueUser         `json:"reporter,omitempty"`
	Watch                         *IssueWatch        `json:"watches,omitempty"`
	UpdateAt                      DateTime           `json:"updated,omitempty"`
	CreatedAt                     DateTime           `json:"created,omitempty"`
	TimeOriginalEstimate          int                `json:"timeoriginalestimate,omitempty"`
	FixVersions                   []*IssueVersion    `json:"fixVersions,omitempty"`
	Epic                          *Epic              `json:"epic,omitempty"`
	Priority                      *IssuePriority     `json:"priority,omitempty"`
	Attachments                   []*IssueAttachment `json:"attachment,omitempty"`
	Assignee                      *IssueUser         `json:"assignee,omitempty"`
	Votes                         *IssueVote         `json:"votes,omitempty"`
	Worklogs                      *IssueWorklogWrap  `json:"worklog,omitempty"`
	DueDate                       DateTime           `json:"duedate,omitempty"`
	Status                        *IssueStatus       `json:"status,omitempty"`
	Creator                       *IssueUser         `json:"creator,omitempty"`
	TimeSpent                     int                `json:"timespent,omitempty"`
	Components                    []*IssueComponent  `json:"components,omitempty"`
	Progress                      *IssueProgress     `json:"progress,omitempty"`
	AggregateProgress             *IssueProgress     `json:"aggregateprogress,omitempty"`
	ResolutionDate                DateTime           `json:"resolutiondate,omitempty"`
	Summary                       string             `json:"summary,omitempty"`
	Comments                      IssueCommentWrap   `json:"comment,omitempty"`
	Versions                      []*IssueVersion    `json:"versions,omitempty"`
}

// IssueType represents the type of Jira Issue
type IssueType struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	SelfLink    string `json:"self,omitempty"`
	Description string `json:"description,omitempty"`
	IconURL     string `json:"iconUrl,omitempty"`
	SubTask     bool   `json:"subtask,omitempty"`
	AvatarID    int    `json:"avatarId,omitempty"`
}

// IssueResolution represents the resolution of Jira Issue
type IssueResolution struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	SelfLink    string `json:"self,omitempty"`
	Description string `json:"description,omitempty"`
}

// IssueLink represents the links of Jira Issue
type IssueLink struct {
	ID       string         `json:"id,omitempty"`
	SelfLink string         `json:"self,omitempty"`
	Type     *IssueLinkType `json:"type,omitempty"`
	Inward   *Issue         `json:"inwardIssue,omitempty"`
	Outward  *Issue         `json:"outwardIssue,omitempty"`
}

// IssueLinkType represents the link type of Jira Issue
type IssueLinkType struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Inward   string `json:"inward,omitempty"`
	Outward  string `json:"outward,omitempty"`
	SelfLink string `json:"self,omitempty"`
}

// IssueTimeTracking represents the time tracking of Jira Issue
type IssueTimeTracking struct {
	OriginalEstimate         string `json:"originalEstimate,omitempty"`
	RemainingEstimate        string `json:"remainingEstimate,omitempty"`
	TimeSpent                string `json:"timeSpent,omitempty"`
	OriginalEstimateSeconds  int    `json:"originalEstimateSeconds,omitempty"`
	RemainingEstimateSeconds int    `json:"remainingEstimateSeconds,omitempty"`
	TimeSpentSeconds         int    `json:"timeSpentSeconds,omitempty"`
}

// IssueUser represents the user of Jira Issue
type IssueUser struct {
	Key         string            `json:"key,omitempty"`
	Name        string            `json:"name,omitempty"`
	SelfLink    string            `json:"self,omitempty"`
	Email       string            `json:"emailAddress,omitempty"`
	DisplayName string            `json:"displayName,omitempty"`
	Active      bool              `json:"active,omitempty"`
	Timezone    string            `json:"timeZone,omitempty"`
	AvatarURLs  map[string]string `json:"avatarUrls,omitempty"`
}

// IssueWatch represents the watch data of Jira Issue
type IssueWatch struct {
	SelfLink string `json:"self,omitempty"`
	Count    int    `json:"watchCount,omitempty"`
	Watching bool   `json:"isWatching,omitempty"`
}

// IssuePriority represents the priority of Jira Issue
type IssuePriority struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	SelfLink string `json:"self,omitempty"`
	IconURL  string `json:"iconUrl,omitempty"`
}

// IssueAttachment represents the attachments list of Jira Issue
type IssueAttachment struct {
	ID        string     `json:"id,omitempty"`
	Filename  string     `json:"filename,omitempty"`
	SelfLink  string     `json:"self,omitempty"`
	Author    *IssueUser `json:"author,omitempty"`
	CreatedAt *DateTime  `json:"created,omitempty"`
	Size      int        `json:"size,omitempty"`
	MimeType  string     `json:"mimeType,omitempty"`
	Content   string     `json:"content,omitempty"`
}

// IssueVote represents the vote data of Jira Issue
type IssueVote struct {
	SelfLink string `json:"self,omitempty"`
	Votes    int    `json:"votes,omitempty"`
	Voted    bool   `json:"hasVoted,omitempty"`
}

// IssueWorklogWrap represents the worklog list of Jira Issue
type IssueWorklogWrap struct {
	Pagination
	Worklogs []*IssueWorklog `json:"worklogs,omitempty"`
}

// IssueWorklog represents the worklog of Jira Issue
type IssueWorklog struct {
	ID               string     `json:"id,omitempty"`
	IssueID          string     `json:"issueId,omitempty"`
	SelfLink         string     `json:"self,omitempty"`
	Author           *IssueUser `json:"author,omitempty"`
	UpdateAuthor     *IssueUser `json:"updateAuthor,omitempty"`
	Comment          string     `json:"comment,omitempty"`
	CreatedAt        DateTime   `json:"created,omitempty"`
	UpdatedAt        DateTime   `json:"updated,omitempty"`
	StartedAt        DateTime   `json:"started,omitempty"`
	TimeSpent        string     `json:"timeSpent,omitempty"`
	TimeSpentSeconds int        `json:"timeSpentSeconds,omitempty"`
}

// IssueStatus represents the status of Jira Issue
type IssueStatus struct {
	ID          string               `json:"id,omitempty"`
	Name        string               `json:"name,omitempty"`
	Description string               `json:"description,omitempty"`
	SelfLink    string               `json:"self,omitempty"`
	IconURL     string               `json:"iconUrl,omitempty"`
	Category    *IssueStatusCategory `json:"statusCategory,omitempty"`
}

// IssueStatusCategory represents the status category of Jira Issue
type IssueStatusCategory struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Key       string `json:"key,omitempty"`
	SelfLink  string `json:"self,omitempty"`
	ColorName string `json:"colorName,omitempty"`
}

// IssueProgress represents the progress of Jira Issue
type IssueProgress struct {
	Progress int `json:"progress,omitempty"`
	Total    int `json:"total,omitempty"`
	Percent  int `json:"percent,omitempty"`
}

// IssueCommentWrap represents the comments list of Jira Issue
type IssueCommentWrap struct {
	Pagination
	Comments []*IssueComment `json:"comments,omitempty"`
}

// IssueComment represents the comment of Jira Issue
type IssueComment struct {
	ID           string    `json:"id,omitempty"`
	SelfLink     string    `json:"self,omitempty"`
	Body         string    `json:"body,omitempty"`
	Author       IssueUser `json:"author,omitempty"`
	UpdateAuthor IssueUser `json:"updateAuthor,omitempty"`
	CreatedAt    DateTime  `json:"created,omitempty"`
	UpdatedAt    DateTime  `json:"updated,omitempty"`
}

// IssueComponent represents the component of Jira Issue
type IssueComponent struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	SelfLink string `json:"self,omitempty"`
}

// IssueVersion represents the version of Jira Issue
type IssueVersion struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	SelfLink    string `json:"self,omitempty"`
	Description string `json:"description,omitempty"`
	Archived    bool   `json:"archived,omitempty"`
	Released    bool   `json:"released,omitempty"`
}

// IssueProjectCategory represents the project category of Jira Issue
type IssueProjectCategory struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	SelfLink    string `json:"self,omitempty"`
	Description string `json:"description,omitempty"`
}

// IssueProject represents a Jira Project
type IssueProject struct {
	ID         string               `json:"id,omitempty"`
	Key        string               `json:"key,omitempty"`
	Name       string               `json:"name,omitempty"`
	SelfLink   string               `json:"self,omitempty"`
	AvatarURLs map[string]string    `json:"avatarUrls,omitempty"`
	Category   IssueProjectCategory `json:"projectCategory,omitempty"`
	Simplified string               `json:"simplified,omitempty"`
	Style      string               `json:"style,omitempty"`
}

// IssuesOptions contains all options to list backlog from a board
type IssuesOptions struct {
	//The starting index of the returned sprints. Base index: 0. See the 'Pagination' section at the top of this page for more details.
	StartAt int `query:"startAt"`
	//The maximum number of sprints to return per page. Default: 50. See the 'Pagination' section at the top of this page for more details.
	MaxResults int `query:"maxResults"`
	//Filters results using a JQL query. If you define an order in your JQL query, it will override the default order of the returned issues.
	JQL string `query:"jql"`
	//Specifies whether to validate the JQL query or not. Default: true.
	ValidateQuery bool `query:"validateQuery"`
	//The list of fields to return for each issue. By default, all navigable and Agile fields are returned.
	Fields string `query:"fields"`
	//This parameter is currently not used.
	Expand string `query:"expand"`
}
