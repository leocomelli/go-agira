package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// IssuesService handles communication with the issues related
// methods of the Jira Agile API
//
// Jira Agile API docs: https://docs.atlassian.com/jira-software/REST/7.3.1/#agile/1.0/issue
type IssuesService service

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
	Project                       *Project           `json:"project,omitempty"`
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

// IssueEstimation represents the estimation of the issue and a fieldId of the field that is used for it
type IssueEstimation struct {
	FieldID string `json:"fieldId,omitempty"`
	Value   int    `json:"value,omitempty"`
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

// GetIssueOptions contains the options to get an issue
type GetIssueOptions struct {
	//The list of fields to return for each issue. By default, all navigable and Agile fields are returned.
	Fields string `query:"fields"`
	//This parameter is currently not used.
	Expand string `query:"expand"`
}

// IssueEstimationOptions contains the options to set the issue estimation
type IssueEstimationOptions struct {
	Value string `json:"value,omitempty"`
}

// Get returns a single issue, for a given issue Id or issue key. Issues returned
// from this resource include Agile fields, like sprint, closedSprints, flagged, and epic.
//
// GET /rest/agile/1.0/issue/{issueIdOrKey}
func (i *IssuesService) Get(ctx context.Context, idOrKey string, opts *GetIssueOptions) (*Issue, *Response, error) {

	q := QueryParameters(opts)

	req, err := i.client.NewRequest("GET", fmt.Sprintf("issue/%s%s", idOrKey, q), nil)
	if err != nil {
		return nil, nil, err
	}

	var issue = &Issue{}
	resp, err := i.client.Do(ctx, req, issue)
	if err != nil {
		return nil, resp, err
	}

	return issue, resp, nil
}

// GetEstimationForBoard returns the estimation of the issue and a fieldId of the field that is used
// for it. boardId param is required. This param determines which field will be updated on a issue.
// Original time internally stores and returns the estimation as a number of seconds.
// The field used for estimation on the given board can be obtained from board configuration resource.
// More information about the field are returned by edit meta resource or field resource.
//
// GET /rest/agile/1.0/issue/{issueIdOrKey}/estimation
func (i *IssuesService) GetEstimationForBoard(ctx context.Context, idOrKey string, boardID int) (*IssueEstimation, *Response, error) {
	req, err := i.client.NewRequest("GET", fmt.Sprintf("issue/%s/estimation?boardId=%d", idOrKey, boardID), nil)
	if err != nil {
		return nil, nil, err
	}

	var issueEst = &IssueEstimation{}
	resp, err := i.client.Do(ctx, req, issueEst)
	if err != nil {
		return nil, resp, err
	}

	return issueEst, resp, nil
}

// EstimationForBoard updates the estimation of the issue. boardId param is required. This param determines
// which field will be updated on a issue.
// Note that this resource changes the estimation field of the issue regardless of appearance the field on the screen.
//
// Original time tracking estimation field accepts estimation in formats like "1w", "2d", "3h", "20m" or number which
// represent number of minutes. However, internally the field stores and returns the estimation as a number of seconds.
//
// The field used for estimation on the given board can be obtained from board configuration resource. More
// information about the field are returned by edit meta resource or field resource.
//
// PUT /rest/agile/1.0/issue/{issueIdOrKey}/estimation
func (i *IssuesService) EstimationForBoard(ctx context.Context, idOrKey string, boardID int, estimation string) (*IssueEstimation, *Response, error) {
	opts := &IssueEstimationOptions{
		Value: estimation,
	}

	req, err := i.client.NewRequest("PUT", fmt.Sprintf("issue/%s/estimation?boardId=%d", idOrKey, boardID), opts)
	if err != nil {
		return nil, nil, err
	}

	var issueEst = &IssueEstimation{}
	resp, err := i.client.Do(ctx, req, issueEst)
	if err != nil {
		return nil, resp, err
	}

	return issueEst, resp, nil
}
