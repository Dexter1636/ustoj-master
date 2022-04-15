package model

type Problem struct {
	ProblemID        int    `gorm:"int"`
	Status           string `gorm:"varchar(20)"`
	Difficulty       string `gorm:"varchar(20)"`
	Acceptance       string `gorm:"varchar(20)"`
	GlobalAcceptance string `gorm:"varchar(20)"`
}

func (Problem) TableName() string {
	return "problem"
}
