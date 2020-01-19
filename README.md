# go-agira

[![Build Status](https://travis-ci.org/leocomelli/go-agira.svg?branch=master)](https://travis-ci.org/leocomelli/go-agira) [![codecov](https://codecov.io/gh/leocomelli/go-agira/branch/master/graph/badge.svg)](https://codecov.io/gh/leocomelli/go-agira)


go-agira is a Go client library for acessing the [Jira Agile API](https://docs.atlassian.com/jira-software/REST/7.3.1/).

## Usage

```go
import "github.com/leocomelli/go-jira"
```

Construct a new Jira Agile client, then use the various services on the client to access different parts of the Jira Agile API. For example:

```go
client, err := jira.NewClient("https://jira.mycompany.com/", nil)
if err != nil {
    // handle error
}

boards, resp, err := client.Boards.ListBoards(context.Background(), nil)
```

Some API methods have optional parameters that can be passed. For example:

```go
client, err := jira.NewClient("https://jira.mycompany.com/", nil)
if err != nil {
    // handle error
}

opts := &jira.ListBoardsOptions{
    ProjectKeyOrID: "CBD",
}

boards, resp, err := client.Boards.ListBoards(context.Background(), opts)
```

### Authentication

The go-jira library does not directly handle authentication. Instead, when creating a new client, pass an http.Client that can handle authentication for you. 

```go
tp := &jira.BasicAuthTransport{
	Username: "myuser",
	Password: "mypass",
}
client, err := jira.NewClient(defaultBaseURL, tp.Client())
if err != nil {
    // handle error
}

// use client
```

### Status

To check the implementation status, [click here](https://github.com/leocomelli/go-agira/blob/master/STATUS.md)
