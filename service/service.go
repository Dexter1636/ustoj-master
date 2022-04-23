package service

import "ustoj-master/model"

type Service interface {
	GetProblemList(map[string]interface{}) ([]*model.Problem, error)
}
