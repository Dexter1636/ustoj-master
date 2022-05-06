package model

type Problem struct {
	ProblemID        int    `gorm:"int" json:"problem_id"`
	Status           string `gorm:"varchar(20)" json:"status"`
	Difficulty       string `gorm:"varchar(20)" json:"difficulty"`
	Acceptance       string `gorm:"varchar(20)" json:"acceptance"`
	GlobalAcceptance string `gorm:"varchar(20)" json:"global_acceptance"`
	ProblemName      string `gorm:"varchar(50)" json:"problem_name"`
}

func (Problem) TableName() string {
	return "problem"
}
