package model

type TestCase struct {
	ProblemId int64  `gorm:"primary_key"`
	CaseId    int    `gorm:"int"`
	Case      string `gorm:"mediumtext"`
	Expected  string `gorm:"mediumtext"`
}

func (TestCase) TableName() string {
	return "test_case"
}
