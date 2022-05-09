package model

import (
	"time"
)

type SubmitJobStatus string

const (
	JobPending SubmitJobStatus = "Pending"
	JobRunning SubmitJobStatus = "Running"
	JobSuccess SubmitJobStatus = "Succeeded"
	JobFailed  SubmitJobStatus = "Failed"
	JobUnknown SubmitJobStatus = "Unknown"
)

type Submission struct {
	SubmissionID   int       `gorm:"primary_key" json:"submission_id"`
	SubmissionTime time.Time `gorm:"timestamp" json:"submission_time"`
	ProblemID      int       `gorm:"bigint" json:"problem_id"`
	Username       string    `gorm:"varchar(20)" json:"username"`
	Language       string    `gorm:"varchar(20)" json:"language"`
	Code           string    `gorm:"mediumtext" json:"code"`
	Status         string    `gorm:"varchar(20)" json:"status"`
	RunTime        int       `gorm:"int" json:"run_time"`
}

func (Submission) TableName() string {
	return "submission"
}
