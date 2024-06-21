package models

type Event struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Organization string  `json:"organization"`
	Date         string  `json:"date"`
	Price        float64 `json:"price"`
	Rating       string  `json:"rating"`
	ImageURL     string  `json:"image_url"`
	CreatedAt    string  `json:"created_at"`
	Location     string  `json:"location"`
}
