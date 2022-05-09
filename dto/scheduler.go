package dto

import "ustoj-master/model"

type SubmissionDto struct {
	SubmissionID int
	ProblemID    int
	Language     string
	Code         string
	Status       model.SubmitJobStatus
}
