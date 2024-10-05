package model


import "github.com/google/uuid"

type User struct{

	ID uuid.UUID `gorm:"type:uuid;primary_key;"`
	
	FirstName string `json:"first_name" gorm:"type:varchar(100);not null"`

	LastName string `json:"last_name" gorm:"type:varchar(100);not null"`

	Email string `json:"email" gorm:"type:varchar(100);unique;not null"`

	Password string `json:"password" gorm:"type:varchar(100);not null"`

	Role string `json:"role" gorm:"type:varchar(100);not null"`

	Product []Product `json:"product" gorm:"foreignKey:UserID"`

}

