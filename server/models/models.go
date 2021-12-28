package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Role     string `gorm:"column:role"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	Token    string `gorm:"column:token"`
}

func (User) TableName() string {
	return "user"
}

type Ticket struct {
	gorm.Model
	ID     uint   `gorm:"primaryKey"`
	FkUser uint   `gorm:"column:fk_user"`
	Title  string `gorm:"column:title"`
	Status string `gorm:"column:status"`
	Body   string `gorm:"column:body"`
}

func (Ticket) TableName() string {
	return "ticket"
}

type Comment struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	FkUser   uint   `gorm:"column:fk_user"`
	FkTicket uint   `gorm:"column:fk_ticket"`
	Body     string `gorm:"column:body"`
}

func (Comment) TableName() string {
	return "comment"
}
