package model

import "github.com/google/uuid"

type Product struct{

	ID uuid.UUID `gorm:"type:uuid;primary_key;"`
	
	Name string `json:"name" gorm:"type:varchar(100);not null"`

	Price float64 `json:"price" gorm:"type:decimal;not null"`

	Quantity int `json:"quantity" gorm:"type:integer;not null"`

	UserID uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`

}


