package model

type Description struct {
	ProblemID   int    `gorm:"primary_key"`
	Description []byte `gorm:"BLOB"`
}

func (Description) TableName() string {
	return "description"
}
