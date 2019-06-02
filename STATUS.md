# Status

[JIRA Agile Server REST API reference v7.3.1](https://docs.atlassian.com/jira-software/REST/7.3.1/#agile/1.0/board-getAllBoards)

## Backlog
* [ ] Move issues to backlog `POST /rest/agile/1.0/backlog/issue`

## Board

* [x] Get all boards `GET /rest/agile/1.0/board`
* [x] Create board `POST /rest/agile/1.0/board`
* [x] Get board `GET /rest/agile/1.0/board/{boardId}`
* [x] Delete board `DELETE /rest/agile/1.0/board/{boardId}`
* [x] Get issues for backlog `GET /rest/agile/1.0/board/{boardId}/backlog`
* [x] Get configuration `GET /rest/agile/1.0/board/{boardId}/configuration`
* [x] Get issues for board `GET /rest/agile/1.0/board/{boardId}/issue`
* [x] Get epics `GET /rest/agile/1.0/board/{boardId}/epic`
* [x] Get issues for epic `GET /rest/agile/1.0/board/{boardId}/epic/{epicId}/issue`
* [x] Get issues without epic `GET /rest/agile/1.0/board/{boardId}/epic/none/issue`
* [x] Get projects `GET /rest/agile/1.0/board/{boardId}/project`
* [ ] Get properties keys `GET /rest/agile/1.0/board/{boardId}/properties`
* [ ] Delete property `DELETE /rest/agile/1.0/board/{boardId}/properties/{propertyKey}`
* [ ] Set property `PUT /rest/agile/1.0/board/{boardId}/properties/{propertyKey}`
* [ ] Get property `GET /rest/agile/1.0/board/{boardId}/properties/{propertyKey}`
* [x] Get all sprints `GET /rest/agile/1.0/board/{boardId}/sprint`
* [x] Get issues for sprint `GET /rest/agile/1.0/board/{boardId}/sprint/{sprintId}/issue`
* [x] Get all versions `GET /rest/agile/1.0/board/{boardId}/version`

## Epic

* [ ] Create epic `POST /rest/api/2/issue`
* [ ] Update issue specific fields of epic `PUT /rest/api/2/issue/{issueIdOrKey}`
* [ ] Delete epic `DELETE /rest/api/2/issue`
* [x] Get epic `GET /rest/agile/1.0/epic/{epicIdOrKey}`
* [x] Partially update epic `POST /rest/agile/1.0/epic/{epicIdOrKey}`
* [x] Get issues for epic `GET /rest/agile/1.0/epic/{epicIdOrKey}/issue`
* [x] Move issues to epic `POST /rest/agile/1.0/epic/{epicIdOrKey}/issue`
* [ ] Rank epics `PUT /rest/agile/1.0/epic/{epicIdOrKey}/rank`
* [x] Get issues without epic `GET /rest/agile/1.0/epic/none/issue`
* [x] Remove issues from epic `POST /rest/agile/1.0/epic/none/issue`

## Issue

* [x] Get issue `GET /rest/agile/1.0/issue/{issueIdOrKey}`
* [x] Get issue estimation for board `GET /rest/agile/1.0/issue/{issueIdOrKey}/estimation`
* [x] Estimate issue for board `PUT /rest/agile/1.0/issue/{issueIdOrKey}/estimation`
* [ ] Rank issues `PUT /rest/agile/1.0/issue/rank`

## Sprint 

* [x] Create sprint `POST /rest/agile/1.0/sprint`
* [x] Get sprint `GET /rest/agile/1.0/sprint/{sprintId}`
* [x] Update sprint `PUT /rest/agile/1.0/sprint/{sprintId}`
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