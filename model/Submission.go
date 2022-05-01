package model

import (
	"time"
)

type Submission struct {
	SubmissionID   int       `gorm:"bigint"`
	SubmissionTime time.Time `gorm:"timestamp"`
	ProblemID      int       `gorm:"bigint"`
	Username       string    `gorm:"varchar(20)"`
	Language       string    `gorm:"varchar(20)"`
	Code           string    `gorm:"mediumtext"`
	Status         string    `gorm:"varchar(20)"`
	RunTime        int       `gorm:"int"`
}

func (Submission) TableName() string {
	return "submission"
}
