package entity

type User struct {
	ID   string `gorm:"primary_key"`
	Name string
}
