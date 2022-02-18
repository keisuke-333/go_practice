package entity

type User struct {
	ID   string `gorm:"primary_key"`
	Name string `gorm:"column:name;type:varchar(64)"`
}
