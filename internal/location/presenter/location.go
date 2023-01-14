package presenter

import "time"

type CreateLocationResponse struct {
	Id        string
	Name      string
	CreatedAt time.Time
	CreatedBy string
	Latitude  float64
	Longitude float64
	Note      string
}

type CreateLocationRequest struct {
	Name      string
	Latitude  float64
	Longitude float64
	Note      string
}
