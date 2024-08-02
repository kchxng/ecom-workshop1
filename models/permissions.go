package models

type Permissions struct {
	Id uint `gorm:"primaryKey; autoIncrement:true" json:"id"`
}
