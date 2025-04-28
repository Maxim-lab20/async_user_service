package model

type User struct {
	ID        uint   `gorm:"id" json:"id"`
	FirstName string `gorm:"column:first_name" json:"first_name"`
	LastName  string `gorm:"column:last_name" json:"last_name"`
}
