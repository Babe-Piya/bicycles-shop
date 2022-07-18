package model

import "time"

type Bicycles struct {
	ID          int       `json:"id"`
	Brand       string    `json:"brand"`
	Model       string    `json:"model"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Buyers struct {
	ID        int       `json:"id"`
	BicycleID int       `json:"bicycle_id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Tel       string    `json:"tel"`
	CreatedAt time.Time `json:"created_at"`
}
