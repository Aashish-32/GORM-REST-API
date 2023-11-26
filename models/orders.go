package models

import "time"

type Order struct {
	Id           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	ProductRefer int     `json:"productid" `
	Product      Product `gorm:"foreignKey:ProductRefer" ` //foreign key will be ProductRefer
	UserRefer    int     `json:"userid" `
	User         User    `gorm:"foreignKey:UserRefer" ` //foreign key will be UserRefer

}
