package models

import (
	"github.com/google/uuid"
)

type Store struct {
	StoreID uuid.UUID `json:"sid"`
	Name    string    `json:"sname"`
}

type Product struct {
	ProductID   int64    `json:"pid"`
	Name        string   `json:"pname"`
	Created_at  string   `json:"created_at"`
	Expiry_date string   `json:"expiry_date"`
	Pmetadata   Metadata `json:"metadata"`
}

type Metadata struct {
	Fat       float64 `json:"fat"`
	Colestrol float64 `json:"colestrol"`
	Protein   float64 `json:"protein"`
}
