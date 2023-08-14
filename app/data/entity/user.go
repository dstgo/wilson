package entity

type User struct {
	UUID     string `gorm:"comment:User UUID;index"`
	Username string `gorm:"comment:username;"`
	Password string `gorm:"comment:User password;"`
	Email    string `gorm:"comment:User concat email;"`
}
