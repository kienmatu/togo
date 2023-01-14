package presenter

import "time"

type TodoResponse struct {
	Id        string
	Content   string
	CreatedAt time.Time
	CreatedBy string
}

type TodoRequest struct {
	Latitude   float64
	Longtitude float64
	Note       string
}
