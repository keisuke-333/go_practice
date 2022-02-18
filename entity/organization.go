package entity

type Organization struct {
	ID   string `gorm:"primary_key"`
	Name string
}
