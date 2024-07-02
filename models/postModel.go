package models

type Post struct {
	Title string `gorm:"not null"`
	Body  string `gorm:"not null"`
	Base  `gorm:"embedded"`
}
