package model

type Description struct {
	ProblemID   int    `gorm:"int"`
	Description []byte `gorm:"BLOB"`
}

func (Description) TableName() string {
	return "description"
}
