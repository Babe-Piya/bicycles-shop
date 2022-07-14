package model

type BicyclesResponse struct {
	Success bool       `json:"success"`
	Data    []Bicycles `json:"data"`
	Message string     `json:"message"`
}
