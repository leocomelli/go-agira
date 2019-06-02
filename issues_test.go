package jira

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var issueAsJSON = `{
       "expand": "renderedFields,names,schema,operations,editmeta,changelog,versionedRepresentations",
       "id": "776509",
       "self": "https://jira.mycompany.com/rest/agile/1.0/issue/776509",
       "key": "MCP-840",
       "fields": {
           "resolution": {
               "self": "https://jira.mycompany.com/rest/api/2/resolution/10000",
               "id": "10000",
               "description": "Ready for Release or associated work has been completed.",
               "name": "Done"
           },
           "lastViewed": "2019-05-07T08:31:01.598+0530",
           "aggregatetimeoriginalestimate": 10800,
           "issuelinks": [
               {
                   "id": "742330",
                   "self": "https://jira.mycompany.com/rest/api/2/issueLink/742330",
                   "type": {
                       "id": "10003",
                       "name": "Relates",
                       "inward": "relates to",
                       "outward": "relates to",
                       "self": "https://jira.mycompany.com/rest/api/2/issueLinkType/10003"
                   },
                   "inwardIssue": {
                       "id": "776510",
                       "key": "MCP-841",
                       "self": "https://jira.mycompany.com/rest/api/2/issue/776510",
                       "fields": {
                           "summary": "summary 1",
                           "status": {
                               "self": "https://jira.mycompany.com/rest/api/2/status/17200",
                               "description": "This status is managed internally by JIRA Software",
                               "iconUrl": "https://jira.mycompany.com/",
                               "name": "Good for the next sprint",
                               "id": "17200",
                               "statusCategory": {
                                   "self": "https://jira.mycompany.com/rest/api/2/statuscategory/4",
                                   "id": 4,
                                   "key": "indeterminate",
                                   "colorName": "yellow",
                                   "name": "In Progress"
                               }
                           },
                           "priority": {
                               "self": "https://jira.mycompany.com/rest/api/2/priority/10002",
                               "iconUrl": "https://jira.mycompany.com/images/icons/priorities/critical.svg",
                               "name": "Critical",
                               "id": "10002"
                           },
                           "issuetype": {
                               "self": "https://jira.mycompany.com/rest/api/2/issuetype/3",
                               "id": "3",
                               "description": "A task that needs to be done.",
                               "iconUrl": "https://jira.mycompany.com/secure/viewavatar?size=xsmall&avatarId=10318&avatarType=issuetype",
                               "name": "Task",
                               "subtask": false,
                               "avatarId": 10318
                           }
                       }
                   }
               }
           ],
           "subtasks":  [
				{
					"id": "903391",
					"key": "MCP-1163",
					"self": "https://jira.mycompany.com/rest/api/2/issue/903391",
					"fields": {
						"summary": "xxxx",
						"status": {
							"self": "https://jira.mycompany.com/rest/api/2/status/10000",
							"description": "This issue is under review.",
							"iconUrl": "https://jira.mycompany.com/images/icons/statuses/open.png",
							"name": "To Do",
							"id": "10000",
							"statusCategory": {
								"self": "https://jira.mycompany.com/rest/api/2/statuscategory/2",
								"id": 2,
								"key": "new",
								"colorName": "blue-gray",
								"name": "To Do"
							}
						},
						"priority": {
							"self": "https://jira.mycompany.com/rest/api/2/priority/10003",
							"iconUrl": "https://jira.mycompany.com/images/icons/priorities/major.svg",
							"name": "Major",
							"id": "10003"
						},
						"issuetype": {
							"self": "https://jira.mycompany.com/rest/api/2/issuetype/5",
							"id": "5",
							"description": "The sub-task of the issue",
							"iconUrl": "https://jira.mycompany.com/secure/viewavatar?size=xsmall&avatarId=10316&avatarType=issuetype",
							"name": "Sub-task",
							"subtask": true,
							"avatarId": 10316
						}
					}
				}
			],
           "closedSprints": [
               {
                   "id": 9666,
                   "self": "https://jira.mycompany.com/rest/agile/1.0/sprint/9666",
                   "state": "closed",
                   "name": "MCP Sprint 17",
                   "startDate": "2019-03-19T16:30:00.000+05:30",
                   "endDate": "2019-03-30T02:30:00.000+05:30",
                   "completeDate": "2019-04-01T22:48:42.603+05:30",
                   "originBoardId": 2881
               },
               {
                   "id": 9963,
                   "self": "https://jira.mycompany.com/rest/agile/1.0/sprint/9963",
                   "state": "closed",
                   "name": "MCP Sprint 18",
                   "startDate": "2019-04-02T22:30:00.000+05:30",
                   "endDate": "2019-04-12T02:30:00.000+05:30",
                   "completeDate": "2019-04-13T00:24:54.648+05:30",
                   "originBoardId": 2881,
                   "goal": ""
               }
           ],
           "customfield_16130": null,
           "issuetype": {
               "self": "https://jira.mycompany.com/rest/api/2/issuetype/1",
               "id": "1",
               "description": "A problem which impairs or prevents the functions of the product.",
               "iconUrl": "https://jira.mycompany.com/secure/viewavatar?size=xsmall&avatarId=10303&avatarType=issuetype",
               "name": "Bug",
               "subtask": false,
               "avatarId": 10303
           },
           "timetracking": {
               "originalEstimate": "3h",
               "remainingEstimate": "1.25h",
               "timeSpent": "5.75h",
               "originalEstimateSeconds": 10800,
               "remainingEstimateSeconds": 4500,
               "timeSpentSeconds": 20700
           },
           "environment": null,
           "timeestimate": 4500,
           "aggregatetimespent": 20700,
           "workratio": 191,
           "flagged": false,
           "labels": [
				"lbl_1",
				"lbl_2"
			],
           "reporter": {
               "self": "https://jira.mycompany.com/rest/api/2/user?username=user1",
               "name": "user1",
               "key": "user1",
               "emailAddress": "user1@mycompany.com",
               "avatarUrls": {
                   "48x48": "https://jira.mycompany.com/secure/useravatar?ownerId=user1&avatarId=22110",
                   "24x24": "https://jira.mycompany.com/secure/useravatar?size=small&ownerId=user1&avatarId=22110",
                   "16x16": "https://jira.mycompany.com/secure/useravatar?size=xsmall&ownerId=user1&avatarId=22110",
                   "32x32": "https://jira.mycompany.com/secure/useravatar?size=medium&ownerId=user1&avatarId=22110"
               },
               "displayName": "User 1",
               "active": true,
               "timeZone": "America/Sao_Paulo"
           },
           "watches": {
               "self": "https://jira.mycompany.com/rest/api/2/issue/MCP-840/watchers",
               "watchCount": 1,
               "isWatching": false
           },
           "updated": "2019-04-03T20:10:49.000+0530",
           "timeoriginalestimate": 10800,
           "description": "my description",
           "fixVersions": [
				{
					"self": "https://jira.mycompany.com/rest/api/2/version/28109",
					"id": "28109",
					"description": "",
					"name": "1.0.0",
					"archived": false,
					"released": false
				}
       	],
           "epic": {
               "id": 540948,
               "key": "MCP-214",
               "self": "https://jira.mycompany.com/rest/agile/1.0/epic/540948",
               "name": "Epic 1",
               "summary": "Epic 1",
               "color": {
                   "key": "color_9"
               },
               "done": false
           },
           "priority": {
               "self": "https://jira.mycompany.com/rest/api/2/priority/10002",
               "iconUrl": "https://jira.mycompany.com/images/icons/priorities/critical.svg",
               "name": "Critical",
               "id": "10002"
           },
           "created": "2019-03-01T02:07:01.000+0530",
           "attachment": [
               {
                   "self": "https://jira.mycompany.com/rest/api/2/attachment/424953",
                   "id": "424953",
                   "filename": "attachment1.txt",
                   "author": {
                       "self": "https://jira.mycompany.com/rest/api/2/user?username=user1",
                       "name": "user1",
                       "key": "user1",
                       "emailAddress": "user1@mycompany.com",
                       "avatarUrls": {
                           "48x48": "https://jira.mycompany.com/secure/useravatar?ownerId=user1&avatarId=22110",
                           "24x24": "https://jira.mycompany.com/secure/useravatar?size=small&ownerId=user1&avatarId=22110",
                           "16x16": "https://jira.mycompany.com/secure/useravatar?size=xsmall&ownerId=user1&avatarId=22110",
                           "32x32": "https://jira.mycompany.com/secure/useravatar?size=medium&ownerId=user1&avatarId=22110"
                       },
                       "displayName": "User 1",
                       "active": true,
                       "timeZone": "America/Sao_Paulo"
                   },
                   "created": "2019-03-01T02:06:18.000+0530",
                   "size": 273920,
                   "mimeType": "application/octet-stream",
                   "content": "https://jira.mycompany.com/secure/attachment/424953/attachment1.txt"
               }
           ],
           "assignee": {
               "self": "https://jira.mycompany.com/rest/api/2/user?username=user2",
               "name": "user2",
               "key": "user2",
               "emailAddress": "user2@mycompany.com",
               "avatarUrls": {
                   "48x48": "https://jira.mycompany.com/secure/useravatar?ownerId=user2&avatarId=20619",
                   "24x24": "https://jira.mycompany.com/secure/useravatar?size=small&ownerId=user2&avatarId=20619",
                   "16x16": "https://jira.mycompany.com/secure/useravatar?size=xsmall&ownerId=user2&avatarId=20619",
                   "32x32": "https://jira.mycompany.com/secure/useravatar?size=medium&ownerId=user2&avatarId=20619"
               },
               "displayName": "User 2",
               "active": true,
               "timeZone": "America/Sao_Paulo"
           },
           "votes": {
               "self": "https://jira.mycompany.com/rest/api/2/issue/MCP-840/votes",
               "votes": 0,
               "hasVoted": false
           },
           "worklog": {
               "startAt": 0,
               "maxResults": 20,
               "total": 4,
               "worklogs": [
                   {
                       "self": "https://jira.mycompany.com/rest/api/2/issue/776509/worklog/359709",
                       "author": {
                           "self": "https://jira.mycompany.com/rest/api/2/user?username=user3",
                           "name": "user3",
                           "key": "user3",
                           "emailAddress": "user3@mycompany.com",
                           "avatarUrls": {
                               "48x48": "https://jira.mycompany.com/secure/useravatar?ownerId=user3&avatarId=20409",
                               "24x24": "https://jira.mycompany.com/secure/useravatar?size=small&ownerId=user3&avatarId=20409",
                               "16x16": "https://jira.mycompany.com/secure/useravatar?size=xsmall&ownerId=user3&avatarId=20409",
                               "32x32": "https://jira.mycompany.com/secure/useravatar?size=medium&ownerId=user3&avatarId=20409"
                           },
                           "displayName": "User 3",
                           "active": true,
                           "timeZone": "America/Sao_Paulo"
                       },
                       "updateAuthor": {
                           "self": "https://jira.mycompany.com/rest/api/2/user?username=user3",
                           "name": "user3",
                           "key": "user3",
                           "emailAddress": "user3@mycompany.com",
                           "avatarUrls": {
                               "48x48": "https://jira.mycompany.com/secure/useravatar?ownerId=user3&avatarId=20409",
                               "24x24": "https://jira.mycompany.com/secure/useravatar?size=small&ownerId=user3&avatarId=20409",
                               "16x16": "https://jira.mycompany.com/secure/useravatar?size=xsmall&ownerId=user3&avatarId=20409",
                               "32x32": "https://jira.mycompany.com/secure/useravatar?size=medium&ownerId=user3&avatarId=20409"
                           },
                           "displayName": "User 3",
                           "active": true,
                           "timeZone": "America/Sao_Paulo"
                       },
                       "comment": "",
                       "created": "2019-03-29T18:42:22.000+0530",
                       "updated": "2019-03-29T18:42:22.000+0530",
                       "started": "2019-03-26T18:42:00.000+0530",
                       "timeSpent": "4h",
                       "timeSpentSeconds": 14400,
                       "id": "359709",
                       "issueId": "776509"
                   },
                   {
                       "self": "https://jira.mycompany.com/rest/api/2/issue/776509/worklog/361254",
                       "author": {
                           "self": "https://jira.mycompany.com/rest/api/2/user?username=user2",
                           "name": "user2",
                           "key": "user2",
                           "emailAddress": "user2@mycompany.com",
                           "avatarUrls": {
                               "48x48": "https://jira.mycompany.com/secure/useravatar?ownerId=user2&avatarId=20619",
                               "24x24": "https://jira.mycompany.com/secure/useravatar?size=small&ownerId=user2&avatarId=20619",
                               "16x16": "https://jira.mycompany.com/secure/useravatar?size=xsmall&ownerId=user2&avatarId=20619",
                               "32x32": "https://jira.mycompany.com/secure/useravatar?size=medium&ownerId=user2&avatarId=20619"
                           },
                           "displayName": "User 2",
                           "active": true,
                           "timeZone": "America/Sao_Paulo"
                       },
                       "updateAuthor": {
                           "self": "https://jira.mycompany.com/rest/api/2/user?username=user2",
                           "name": "user2",
                           "key": "user2",
                           "emailAddress": "user2@mycompany.com",
                           "avatarUrls": {
                               "48x48": "https://jira.mycompany.com/secure/useravatar?ownerId=user2&avatarId=20619",
                               "24x24": "https://jira.mycompany.com/secure/useravatar?size=small&ownerId=user2&avatarId=20619",
                               "16x16": "https://jira.mycompany.com/secure/useravatar?size=xsmall&ownerId=user2&avatarId=20619",
                               "32x32": "https://jira.mycompany.com/secure/useravatar?size=medium&ownerId=user2&avatarId=20619"
                           },
                           "displayName": "User 2",
                           "active": true,
                           "timeZone": "America/Sao_Paulo"
                       },
                       "comment": "",
                       "created": "2019-04-02T23:02:05.000+0530",
                       "updated": "2019-04-02T23:02:05.000+0530",
                       "started": "2019-04-02T23:02:00.000+0530",
                       "timeSpent": "1h",
                       "timeSpentSeconds": 3600,
                       "id": "361254",
                       "issueId": "776509"
                   },
                   {
                       "self": "https://jira.mycompany.com/rest/api/2/issue/776509/worklog/361271",
                       "author": {
                           "self": "https://jira.mycompany.com/rest/api/2/user?username=user2",
                           "name": "user2",
                           "key": "user2",
                           "emailAddress": "user2@mycompany.com",
                           "avatarUrls": {
                               "48x48": "https://jira.mycompany.com/secure/useravatar?ownerId=user2&avatarId=20619",
                               "24x24": "https://jira.mycompany.com/secure/useravatar?size=small&ownerId=user2&avatarId=20619",
                               "16x16": "https://jira.mycompany.com/secure/useravatar?size=xsmall&ownerId=user2&avatarId=20619",
                               "32x32": "https://jira.mycompany.com/secure/useravatar?size=medium&ownerId=user2&avatarId=20619"
                           },
                           "displayName": "User 2",
                           "active": true,
                           "timeZone": "America/Sao_Paulo"
                       },
                       "updateAuthor": {
                           "self": "https://jira.mycompany.com/rest/api/2/user?username=user2",
                           "name": "user2",
                           "key": "user2",
                           "emailAddress": "user2@mycompany.com",
                           "avatarUrls": {
                               "48x48": "https://jira.mycompany.com/secure/useravatar?ownerId=user2&avatarId=20619",
                               "24x24": "https://jira.mycompany.com/secure/useravatar?size=small&ownerId=user2&avatarId=20619",
                               "16x16": "https://jira.mycompany.com/secure/useravatar?size=xsmall&ownerId=user2&avatarId=20619",
                               "32x32": "https://jira.mycompany.com/secure/useravatar?size=medium&ownerId=user2&avatarId=20619"
                           },
                           "displayName": "User 2",
                           "active": true,
                           "timeZone": "America/Sao_Paulo"
                       },
                       "comment": "",
                       "created": "2019-04-03T01:14:36.000+0530",
                       "updated": "2019-04-03T01:14:36.000+0530",
                       "started": "2019-04-03T01:14:00.000+0530",
                       "timeSpent": "30m",
                       "timeSpentSeconds": 1800,
                       "id": "361271",
                       "issueId": "776509"
                   },
                   {
                       "self": "https://jira.mycompany.com/rest/api/2/issue/776509/worklog/361882",
                       "author": {
                           "self": "https://jira.mycompany.com/rest/api/2/user?username=user4",
                           "name": "user4",
                           "key": "user4",
                           "emailAddress": "user4@mycompany.com",
                           "avatarUrls": {
                               "48x48": "https://jira.mycompany.com/secure/useravatar?ownerId=user4&avatarId=21302",
                               "24x24": "https://jira.mycompany.com/secure/useravatar?size=small&ownerId=user4&avatarId=21302",
                               "16x16": "https://jira.mycompany.com/secure/useravatar?size=xsmall&ownerId=user4&avatarId=21302",
                               "32x32": "https://jira.mycompany.com/secure/useravatar?size=medium&ownerId=user4&avatarId=21302"
                           },
                           "displayName": "User 4 ",
                           "active": true,
                           "timeZone": "Asia/Calcutta"
                       },
                       "updateAuthor": {
                           "self": "https://jira.mycompany.com/rest/api/2/user?username=user4",
                           "name": "user4",
                           "key": "user4",
                           "emailAddress": "user4@mycompany.com",
                           "avatarUrls": {
                               "48x48": "https://jira.mycompany.com/secure/useravatar?ownerId=user4&avatarId=21302",
                               "24x24": "https://jira.mycompany.com/secure/useravatar?size=small&ownerId=user4&avatarId=21302",
                               "16x16": "https://jira.mycompany.com/secure/useravatar?size=xsmall&ownerId=user4&avatarId=21302",
                               "32x32": "https://jira.mycompany.com/secure/useravatar?size=medium&ownerId=user4&avatarId=21302"
                           },
                           "displayName": "User 4 ",
                           "active": true,
                           "timeZone": "Asia/Calcutta"
                       },
                       "comment": "",
                       "created": "2019-04-03T20:10:43.000+0530",
                       "updated": "2019-04-03T20:10:43.000+0530",
                       "started": "2019-04-03T20:10:00.000+0530",
                       "timeSpent": "15m",
                       "timeSpentSeconds": 900,
                       "id": "361882",
                       "issueId": "776509"
                   }
               ]
           },
           "duedate": null,
           "status": {
               "self": "https://jira.mycompany.com/rest/api/2/status/10001",
               "description": "Work has been completed on this issue.",
               "iconUrl": "https://jira.mycompany.com/images/icons/statuses/closed.png",
               "name": "Done",
               "id": "10001",
               "statusCategory": {
                   "self": "https://jira.mycompany.com/rest/api/2/statuscategory/3",
                   "id": 3,
                   "key": "done",
                   "colorName": "green",
                   "name": "Done"
               }
           },
           "aggregatetimeestimate": 4500,
           "creator": {
               "self": "https://jira.mycompany.com/rest/api/2/user?username=user1",
               "name": "user1",
               "key": "user1",
               "emailAddress": "user1@mycompany.com",
               "avatarUrls": {
                   "48x48": "https://jira.mycompany.com/secure/useravatar?ownerId=user1&avatarId=22110",
                   "24x24": "https://jira.mycompany.com/secure/useravatar?size=small&ownerId=user1&avatarId=22110",
                   "16x16": "https://jira.mycompany.com/secure/useravatar?size=xsmall&ownerId=user1&avatarId=22110",
                   "32x32": "https://jira.mycompany.com/secure/useravatar?size=medium&ownerId=user1&avatarId=22110"
               },
               "displayName": "User 1",
               "active": true,
               "timeZone": "America/Sao_Paulo"
           },
           "timespent": 20700,
           "components": [
               {
                   "self": "https://jira.mycompany.com/rest/api/2/component/25104",
                   "id": "25104",
                   "name": "Cmp 1"
               },
               {
                   "self": "https://jira.mycompany.com/rest/api/2/component/25103",
                   "id": "25103",
                   "name": "Cmp 2"
               }
           ],
           "progress": {
               "progress": 20700,
               "total": 25200,
               "percent": 82
           },
           "project": {
               "self": "https://jira.mycompany.com/rest/api/2/project/17526",
               "id": "17526",
               "key": "MCP",
               "name": "Project 1",
               "avatarUrls": {
                   "48x48": "https://jira.mycompany.com/secure/projectavatar?pid=17526&avatarId=20500",
                   "24x24": "https://jira.mycompany.com/secure/projectavatar?size=small&pid=17526&avatarId=20500",
                   "16x16": "https://jira.mycompany.com/secure/projectavatar?size=xsmall&pid=17526&avatarId=20500",
                   "32x32": "https://jira.mycompany.com/secure/projectavatar?size=medium&pid=17526&avatarId=20500"
               }
           },
           "resolutiondate": "2019-04-03T20:10:49.000+0530",
           "summary": "Summary Xxxx Yyyyy ",
           "comment": {
           "comments": [
               {
                   "self": "https://jira.mycompany.com/rest/api/2/issue/864187/comment/1309127",
                   "id": "1309127",
                   "author": {
                       "self": "https://jira.mycompany.com/rest/api/2/user?username=user3",
                       "name": "user3",
                       "key": "user3",
                       "emailAddress": "user3@mycompany.com",
                       "avatarUrls": {
                           "48x48": "https://jira.mycompany.com/secure/useravatar?ownerId=user3&avatarId=20409",
                           "24x24": "https://jira.mycompany.com/secure/useravatar?size=small&ownerId=user3&avatarId=20409",
                           "16x16": "https://jira.mycompany.com/secure/useravatar?size=xsmall&ownerId=user3&avatarId=20409",
                           "32x32": "https://jira.mycompany.com/secure/useravatar?size=medium&ownerId=user3&avatarId=20409"
                       },
                       "displayName": "User 3",
                       "active": true,
                       "timeZone": "America/Sao_Paulo"
                   },
                   "body": "comment body",
                   "updateAuthor": {
                       "self": "https://jira.mycompany.com/rest/api/2/user?username=user3",
                       "name": "user3",
                       "key": "user3",
                       "emailAddress": "user3@mycompany.com",
                       "avatarUrls": {
                           "48x48": "https://jira.mycompany.com/secure/useravatar?ownerId=user3&avatarId=20409",
                           "24x24": "https://jira.mycompany.com/secure/useravatar?size=small&ownerId=user3&avatarId=20409",
                           "16x16": "https://jira.mycompany.com/secure/useravatar?size=xsmall&ownerId=user3&avatarId=20409",
                           "32x32": "https://jira.mycompany.com/secure/useravatar?size=medium&ownerId=user3&avatarId=20409"
                       },
                       "displayName": "User 3",
                       "active": true,
                       "timeZone": "America/Sao_Paulo"
                   },
                   "created": "2019-05-22T02:04:43.000+0530",
                   "updated": "2019-05-22T02:04:43.000+0530"
               },
               {
                   "self": "https://jira.mycompany.com/rest/api/2/issue/864187/comment/1311276",
                   "id": "1311276",
                   "author": {
                       "self": "https://jira.mycompany.com/rest/api/2/user?username=user3",
                       "name": "user3",
                       "key": "user3",
                       "emailAddress": "user3@mycompany.com",
                       "avatarUrls": {
                           "48x48": "https://jira.mycompany.com/secure/useravatar?ownerId=user3&avatarId=20409",
                           "24x24": "https://jira.mycompany.com/secure/useravatar?size=small&ownerId=user3&avatarId=20409",
                           "16x16": "https://jira.mycompany.com/secure/useravatar?size=xsmall&ownerId=user3&avatarId=20409",
                           "32x32": "https://jira.mycompany.com/secure/useravatar?size=medium&ownerId=user3&avatarId=20409"
                       },
                       "displayName": "User 3",
                       "active": true,
                       "timeZone": "America/Sao_Paulo"
                   },
                   "body": "[~user5] verifique se as soluções acima atendem as nossas necessidades.",
                   "updateAuthor": {
                       "self": "https://jira.mycompany.com/rest/api/2/user?username=user3",
                       "name": "user3",
                       "key": "user3",
                       "emailAddress": "user3@mycompany.com",
                       "avatarUrls": {
                           "48x48": "https://jira.mycompany.com/secure/useravatar?ownerId=user3&avatarId=20409",
                           "24x24": "https://jira.mycompany.com/secure/useravatar?size=small&ownerId=user3&avatarId=20409",
                           "16x16": "https://jira.mycompany.com/secure/useravatar?size=xsmall&ownerId=user3&avatarId=20409",
                           "32x32": "https://jira.mycompany.com/secure/useravatar?size=medium&ownerId=user3&avatarId=20409"
                       },
                       "displayName": "User 3",
                       "active": true,
                       "timeZone": "America/Sao_Paulo"
                   },
                   "created": "2019-05-22T20:23:37.000+0530",
                   "updated": "2019-05-22T20:23:37.000+0530"
               }
           ],
           "maxResults": 2,
           "total": 2,
           "startAt": 0
       	},
           "versions": [
				{
					"self": "https://jira.mycompany.com/rest/api/2/version/28109",
					"id": "28109",
					"description": "",
					"name": "1.0.0",
					"archived": false,
					"released": false
				}
       	],
           "aggregateprogress": {
               "progress": 20700,
               "total": 25200,
               "percent": 82
           }
       }
   }`

var issuesAsJSON = fmt.Sprintf(`{
   "expand": "schema,names",
   "startAt": 0,
   "maxResults": 50,
   "total": 13,
   "issues": [%s]
}`, issueAsJSON)

func TestIssuesServiceGet(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/issue/5", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		fmt.Fprint(w, issueAsJSON)
	})

	issue, _, err := client.Issues.Get(context.Background(), "5", &GetIssueOptions{})
	assert.Nil(t, err)

	assert.NotNil(t, issue)
	assert.Equal(t, "Project 1", issue.Fields.Project.Name)
}
