package models

type User struct {
	Id       int64  `gorm: "primaryKey" json:"id"`
	Email    string `gorm: "varchar(300)" json:"email"`
	FullName string `gorm: "varchar(300)" json:"fullName"`
	Phone    string `gorm: "varchar(300)" json:"phone"`
	Photo    string `gorm: "varchar(300)" json:"photo"`
	Password string `gorm: "varchar(300)" json:"password"`
}
