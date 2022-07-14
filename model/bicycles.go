package model

import "time"

type Bicycles struct {
	ID          int       `json:"id"`
	Brand       string    `json:"brand"`
	Model       string    `json:"model"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
