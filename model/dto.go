package model

const (
	SELL = 1
	BUY  = 2
)

type BicyclesResponse struct {
	Success bool       `json:"success"`
	Data    []Bicycles `json:"data"`
	Message string     `json:"message"`
}

type BuyBicycleResponse struct {
	Success bool        `json:"success"`
	Data    BicycleData `json:"data"`
	Message string      `json:"message"`
}

type BicycleData struct {
	ID    int     `json:"id"`
	Brand string  `json:"brand"`
	Model string  `json:"model"`
	Price float64 `json:"price"`
}

type BuyBicycleRequest struct {
	BicycleID int    `json:"bicycle_id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Tel       string `json:"tel"`
}

type BicycleRequest struct {
	Brand       string  `json:"brand"`
	Model       string  `json:"model"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type CreateBicycleResponse struct {
	Success bool     `json:"success"`
	Data    Bicycles `json:"data"`
	Message string   `json:"message"`
}
