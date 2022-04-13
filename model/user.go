package model

type User struct {
	Username string `gorm:"varchar(20)"`
	Password string `gorm:"varchar(20)"`
	RoleId   int8   `gorm:"tinyint"`
}

func (User) TableName() string {
	return "user"
}
