package entity

import "time"

type Cat struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name"`
	Breed     string    `json:"breed"`
	Color     string    `json:"color"`
	BirthAt   time.Time `json:"birthAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
