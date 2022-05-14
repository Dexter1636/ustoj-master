package model

type User struct {
	ID       int64  `gorm:"column:Id;primary_key"                                    json:"id"`
	Username string `gorm:"varchar(20)"`
	Password string `gorm:"varchar(20)"`
	RoleId   int8   `gorm:"tinyint"`
}

func (User) TableName() string {
	return "user"
}
