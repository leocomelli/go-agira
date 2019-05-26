# Status

[JIRA Agile Server REST API reference v7.3.1](https://docs.atlassian.com/jira-software/REST/7.3.1/#agile/1.0/board-getAllBoards)

## Backlog
* [ ] Move issues to backlog `POST /rest/agile/1.0/backlog/issue`

## Board

* [x] Get all boards `GET /rest/agile/1.0/board`
* [ ] Create board `POST /rest/agile/1.0/board`
* [x] Get board `GET /rest/agile/1.0/board/{boardId}`
* [ ] Delete board `DELETE /rest/agile/1.0/board/{boardId}`
* [x] Get issues for backlog `GET /rest/agile/1.0/board/{boardId}/backlog`
* [ ] Get configuration `GET /rest/agile/1.0/board/{boardId}/configuration`
* [ ] Get issues for board `GET /rest/agile/1.0/board/{boardId}/issue`
* [x] Get epics `GET /rest/agile/1.0/board/{boardId}/epic`
* [ ] Get issues for epic `GET /rest/agile/1.0/board/{boardId}/epic/{epicId}/issue`
* [ ] Get issues without epic `GET /rest/agile/1.0/board/{boardId}/epic/none/issue`
* [ ] Get projects `GET /rest/agile/1.0/board/{boardId}/project`
* [ ] Get properties keys `GET /rest/agile/1.0/board/{boardId}/properties`
* [ ] Delete property `DELETE /rest/agile/1.0/board/{boardId}/properties/{propertyKey}`
* [ ] Set property `PUT /rest/agile/1.0/board/{boardId}/properties/{propertyKey}`
* [ ] Get property `GET /rest/agile/1.0/board/{boardId}/properties/{propertyKey}`
* [x] Get all sprints `GET /rest/agile/1.0/board/{boardId}/sprint`
* [ ] Get issues for sprint `GET /rest/agile/1.0/board/{boardId}/sprint/{sprintId}/issue`
* [ ] Get all versions `GET /rest/agile/1.0/board/{boardId}/version`

## Epic

* [ ] Create epic `POST /rest/api/2/issue`
* [ ] Update issue specific fields of epic `PUT /rest/api/2/issue/{issueIdOrKey}`
* [ ] Delete epic `DELETE /rest/api/2/issue`
* [ ] Get epic `GET /rest/agile/1.0/epic/{epicIdOrKey}`
* [ ] Partially update epic `POST /rest/agile/1.0/epic/{epicIdOrKey}`
* [ ] Get issues for epic `GET /rest/agile/1.0/epic/{epicIdOrKey}/issue`
* [ ] Move issues to epic `POST /rest/agile/1.0/epic/{epicIdOrKey}/issue`
* [ ] Rank epics `PUT /rest/agile/1.0/epic/{epicIdOrKey}/rank`
* [ ] Get issues without epic `GET /rest/agile/1.0/epic/none/issue`
* [ ] Remove issues from epic `POST /rest/agile/1.0/epic/none/issue`

## Issue

* [ ] Get issue `GET /rest/agile/1.0/issue/{issueIdOrKey}`
* [ ] Get issue estimation for board `GET /rest/agile/1.0/issue/{issueIdOrKey}/estimation`
* [ ] Estimate issue for board `PUT /rest/agile/1.0/issue/{issueIdOrKey}/estimation`
* [ ] Rank issues `PUT /rest/agile/1.0/issue/rank`

## Sprint 

* [ ] Create sprint `POST /rest/agile/1.0/sprint`
* [ ] Get sprint `GET /rest/agile/1.0/sprint/{sprintId}`
* [ ] Update sprint `PUT /rest/agile/1.0/sprint/{sprintId}`
* [ ] Partially update sprint `POST /rest/agile/1.0/sprint/{sprintId}`
* [ ] Delete sprint `DELETE /rest/agile/1.0/sprint/{sprintId}`
* [ ] Move issues to sprint `POST /rest/agile/1.0/sprint/{sprintId}/issue`
* [ ] Get issues for sprint `GET /rest/agile/1.0/sprint/{sprintId}/issue`
* [ ] Swap sprint `POST /rest/agile/1.0/sprint/{sprintId}/swap`

## Properties

* [ ] Get properties keys `GET /rest/agile/1.0/sprint/{sprintId}/properties`
* [ ] Delete property `DELETE /rest/agile/1.0/sprint/{sprintId}/properties/{propertyKey}`
* [ ] Set property `PUT /rest/agile/1.0/sprint/{sprintId}/properties/{propertyKey}`
* [ ] Get property `GET /rest/agile/1.0/sprint/{sprintId}/properties/{propertyKey}`