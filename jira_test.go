package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	defaultBaseURL = "https://jira.com/"
	baseURLPath    = "/agile/1.0"
)

type User struct {
	Login string `json:"login"`
}

func setup() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	mux = http.NewServeMux()

	apiHandler := http.NewServeMux()
	apiHandler.Handle(baseURLPath+"/", http.StripPrefix(baseURLPath, mux))
	apiHandler.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(os.Stderr, "FAIL: Client.BaseURL path prefix is not preserved in the request URL:")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "\t"+req.URL.String())
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "\tDid you accidentally use an absolute endpoint URL rather than relative?")
		http.Error(w, "Client.BaseURL path prefix is not preserved in the request URL.", http.StatusInternalServerError)
	})

	server := httptest.NewServer(apiHandler)

	client, _ = NewClient(defaultBaseURL, nil)
	url, _ := url.Parse(server.URL + baseURLPath + "/")
	client.BaseURL = url

	return client, mux, server.URL, server.Close
}

func TestNewRequest(t *testing.T) {
	c, _ := NewClient(defaultBaseURL, nil)

	inURL, outURL := "/foo", defaultBaseURL+"foo"
	inBody, outBody := &User{Login: "l"}, `{"login":"l"}`+"\n"
	req, _ := c.NewRequest("GET", inURL, inBody)

	assert.Equal(t, outURL, req.URL.String())

	body, _ := ioutil.ReadAll(req.Body)
	assert.Equal(t, outBody, string(body))
}

func TestNewRequestInvalidJSON(t *testing.T) {
	c, _ := NewClient(defaultBaseURL, nil)

	type T struct {
		A map[interface{}]interface{}
	}
	_, err := c.NewRequest("GET", ".", &T{})

	assert.NotNil(t, err)

	_, ok := err.(*json.UnsupportedTypeError)
	assert.True(t, ok)
}

func TestNewRequestBadURL(t *testing.T) {
	c, _ := NewClient(defaultBaseURL, nil)

	_, err := c.NewRequest("GET", ":", nil)

	assert.NotNil(t, err)

	uerr, ok := err.(*url.Error)
	assert.True(t, ok)
	assert.Equal(t, "parse", uerr.Op)
}

func TestNewRequestEmptyBody(t *testing.T) {
	c, _ := NewClient(defaultBaseURL, nil)
	req, err := c.NewRequest("GET", ".", nil)

	assert.Nil(t, err)
	assert.Nil(t, req.Body)
}

func TestNewRequestErrorForNoTrailingSlash(t *testing.T) {
	tests := []struct {
		rawurl    string
		wantError bool
	}{
		{rawurl: "https://jira.com", wantError: true},
		{rawurl: "https://jira.com/", wantError: false},
	}
	c, _ := NewClient(defaultBaseURL, nil)
	for _, test := range tests {
		u, err := url.Parse(test.rawurl)
		assert.Nil(t, err)

		c.BaseURL = u
		_, err = c.NewRequest(http.MethodGet, "test", nil)

		if test.wantError {
			assert.NotNil(t, err)
		} else if !test.wantError {
			assert.Nil(t, err)
		}
	}
}

func TestDo(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		fmt.Fprint(w, `{"login":"foo"}`)
	})

	req, _ := client.NewRequest("GET", ".", nil)
	body := &User{}
	want := &User{"foo"}
	client.Do(context.Background(), req, body)

	assert.True(t, reflect.DeepEqual(body, want))
}

func TestDoHTTPError(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	req, _ := client.NewRequest("GET", ".", nil)
	resp, err := client.Do(context.Background(), req, nil)

	assert.NotNil(t, err)
	assert.Equal(t, 400, resp.StatusCode)
}

func TestDoNoContent(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	var body json.RawMessage

	req, _ := client.NewRequest("GET", ".", nil)
	_, err := client.Do(context.Background(), req, &body)

	assert.Nil(t, err)
}

func TestBasicAuthTransport(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	username, password := "u", "p"

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()

		assert.True(t, ok)
		assert.Equal(t, username, u)
		assert.Equal(t, password, p)
	})

	tp := &BasicAuthTransport{
		Username: username,
		Password: password,
	}
	bac, _ := NewClient(defaultBaseURL, tp.Client())
	bac.BaseURL = client.BaseURL
	req, _ := bac.NewRequest("GET", ".", nil)
	bac.Do(context.Background(), req, nil)
}

func TestQueryParameters(t *testing.T) {

	type MyOptions struct {
		MaxResults int    `query:"maxResults"`
		Name       string `query:"name"`
		IsLast     bool   `query:"isLast"`
	}

	tests := []struct {
		Name      string
		Options   *MyOptions
		Query     string
		Assetions int
	}{
		{
			Name: "all options",
			Options: &MyOptions{
				MaxResults: 50,
				Name:       "foo",
				IsLast:     true,
			},
			Query:     "?maxResults=50&name=foo&isLast=true",
			Assetions: 3,
		},
		{
			Name: "one options",
			Options: &MyOptions{
				Name: "foo",
			},
			Query:     "?name=foo",
			Assetions: 1,
		},
		{
			Name:      "empty options",
			Options:   &MyOptions{},
			Query:     "",
			Assetions: 0,
		},
		{
			Name:      "nil options",
			Options:   nil,
			Query:     "",
			Assetions: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			s := QueryParameters(tt.Options)

			vars := strings.Split(tt.Query, "&")
			assertions := 0
			for _, v := range vars {
				if tt.Query != "" && strings.Contains(s, strings.TrimPrefix(v, "?")) {
					assertions++
				}
			}

			assert.Equal(t, tt.Assetions, assertions)
		})
	}
}
