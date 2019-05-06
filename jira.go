package jira

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/fatih/structs"
)

// A Client manages communication with the Jira Agile API.
type Client struct {
	client  *http.Client
	BaseURL *url.URL
	Path    string

	// Reuse a single struct instead of allocating one for each service on the heap.
	common service

	Boards *BoardsService
}

type service struct {
	client *Client
}

// NewClient returns a new Jira Agile API client. If a nil httpClient is
// provided, http.DefaultClient will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).
func NewClient(baseURL string, httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseEndpoint, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	if !strings.HasSuffix(baseEndpoint.Path, "/") {
		baseEndpoint.Path += "/"
	}

	c := &Client{
		client:  httpClient,
		BaseURL: baseEndpoint,
	}
	c.common.client = c
	c.Boards = (*BoardsService)(&c.common)

	return c, nil
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}
	defer resp.Body.Close()

	response := &Response{
		Response: resp,
	}

	if code := resp.StatusCode; code < 200 || code > 299 {
		errResp := &ErrorResponse{
			Response: resp,
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err == nil && data != nil {
			json.Unmarshal(data, errResp)
		}
		return response, errResp
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return response, err
}

// Pagination contains the information about pagination
type Pagination struct {
	MaxResults int  `json:"maxResults,omitempty"`
	StartAt    int  `json:"startAt,omitempty"`
	IsLast     bool `json:"isLast,omitempty"`
}

// Response is a Jira Agile API response. This wraps the standard http.Response
// returned from Jira and provides convenient access to things like
// pagination info.
type Response struct {
	*http.Response
	Pagination
}

// ErrorResponse reports one or more errors caused by an API request.
type ErrorResponse struct {
	Response *http.Response
	Messages []string          `json:"errorMessages,omitempty"`
	Errors   map[string]string `json:"errors,omitempty"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %+v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Messages, r.Errors)
}

// BasicAuthTransport is an http.RoundTripper that authenticates all requests
// using HTTP Basic Authentication with the provided username and password.
type BasicAuthTransport struct {
	Transport http.RoundTripper
	Username  string
	Password  string
}

// RoundTrip implements the RoundTripper interface.
func (t *BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// To set extra headers, we must make a copy of the Request so
	// that we don't modify the Request we were given. This is required by the
	// specification of http.RoundTripper.
	//
	// Since we are going to modify only req.Header here, we only need a deep copy
	// of req.Header.
	req2 := new(http.Request)
	*req2 = *req
	req2.Header = make(http.Header, len(req.Header))
	for k, s := range req.Header {
		req2.Header[k] = append([]string(nil), s...)
	}

	req2.SetBasicAuth(t.Username, t.Password)

	if t.Transport != nil {
		return t.Transport.RoundTrip(req2)
	}
	return http.DefaultTransport.RoundTrip(req2)

}

// Client returns an *http.Client that makes requests that are authenticated
// using HTTP Basic Authentication.
func (t *BasicAuthTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

// QueryParameters returns a query parameters string to use in the request.
// Some endpoint allow options using query parameters, this method returns a
// string as expected: ?k1=v1&k2=v2&k3=v3
func QueryParameters(val interface{}) string {
	if val == nil || (reflect.ValueOf(val).Kind() == reflect.Ptr && reflect.ValueOf(val).IsNil()) {
		return ""
	}

	var query []string

	s := structs.New(val)
	m := s.Map()

	for k, v := range m {
		f := s.Field(k)
		t := f.Tag("query")

		if !f.IsZero() {
			query = append(query, fmt.Sprintf("%v=%v", t, v))
		}
	}

	if len(query) == 0 {
		return ""
	}

	return "?" + strings.Join(query, "&")
}
