package models

type Data struct {
	Events []Event `json:"events"`
	Spots  []Spot  `json:"spots"`
}
