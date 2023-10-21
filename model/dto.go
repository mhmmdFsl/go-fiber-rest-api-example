package model

import "time"

type AddCatDto struct {
	Name    string    `json:"name"`
	Breed   string    `json:"breed"`
	Color   string    `json:"color"`
	BirthAt time.Time `json:"birthAt"`
}
