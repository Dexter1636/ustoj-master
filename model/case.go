package model

type Case struct {
	ProblemId int64  `gorm:"bigint"`
	CaseId    int    `gorm:"int"`
	Case      string `gorm:"mediumtext"`
	Expected  string `gorm:"mediumtext"`
}

func (Case) TableName() string {
	return "case"
}
