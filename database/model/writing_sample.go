package dbmodel

import "gorm.io/gorm"

type WritingSample struct {
	gorm.Model
	ID      string `gorm:"unique;default:gen_random_uuid()"`
	Title   string
	Content string
	UserID  string
}
