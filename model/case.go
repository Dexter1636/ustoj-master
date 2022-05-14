package model

type TestCase struct {
	ProblemId int64  `gorm:"int"`
	CaseId    int    `gorm:"primary_key"`
	Case      string `gorm:"mediumtext"`
	Expected  string `gorm:"mediumtext"`
}

func (TestCase) TableName() string {
	return "test_case"
}
