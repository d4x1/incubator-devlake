/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package migrationscripts

import (
	"context"
	"time"

	"github.com/apache/incubator-devlake/models/migrationscripts/archived"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type JiraIssue20220507 struct {
	// collected fields
	SourceId                 uint64 `gorm:"primaryKey"`
	IssueId                  uint64 `gorm:"primarykey"`
	ProjectId                uint64
	Self                     string `gorm:"type:varchar(255)"`
	IconURL                  string `gorm:"type:varchar(255);column:icon_url"`
	Key                      string `gorm:"type:varchar(255)"`
	Summary                  string
	Type                     string `gorm:"type:varchar(255)"`
	EpicKey                  string `gorm:"type:varchar(255)"`
	StatusName               string `gorm:"type:varchar(255)"`
	StatusKey                string `gorm:"type:varchar(255)"`
	StoryPoint               float64
	OriginalEstimateMinutes  int64  // user input?
	AggregateEstimateMinutes int64  // sum up of all subtasks?
	RemainingEstimateMinutes int64  // could it be negative value?
	CreatorAccountId         string `gorm:"type:varchar(255)"`
	CreatorAccountType       string `gorm:"type:varchar(255)"`
	CreatorDisplayName       string `gorm:"type:varchar(255)"`
	AssigneeAccountId        string `gorm:"type:varchar(255);comment:latest assignee"`
	AssigneeAccountType      string `gorm:"type:varchar(255)"`
	AssigneeDisplayName      string `gorm:"type:varchar(255)"`
	PriorityId               uint64
	PriorityName             string `gorm:"type:varchar(255)"`
	ParentId                 uint64
	ParentKey                string `gorm:"type:varchar(255)"`
	SprintId                 uint64 // latest sprint, issue might cross multiple sprints, would be addressed by #514
	SprintName               string `gorm:"type:varchar(255)"`
	ResolutionDate           *time.Time
	Created                  time.Time
	Updated                  time.Time `gorm:"index"`
	SpentMinutes             int64
	LeadTimeMinutes          uint
	StdStoryPoint            uint
	StdType                  string `gorm:"type:varchar(255)"`
	StdStatus                string `gorm:"type:varchar(255)"`
	AllFields                datatypes.JSONMap

	// internal status tracking
	ChangelogUpdated  *time.Time
	RemotelinkUpdated *time.Time
	archived.NoPKModel
}

func (JiraIssue20220507) TableName() string {
	return "_tool_jira_issues"
}

type UpdateSchemas20220507 struct{}

func (*UpdateSchemas20220507) Up(ctx context.Context, db *gorm.DB) error {
	err := db.Migrator().AddColumn(&JiraIssue20220507{}, "icon_url")
	if err != nil {
		return err
	}
	return nil
}

func (*UpdateSchemas20220507) Version() uint64 {
	return 20220507154646
}

func (*UpdateSchemas20220507) Name() string {
	return "Add icon_url column to JiraIssue"
}