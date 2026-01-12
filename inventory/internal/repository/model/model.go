package model

import "time"

type Part struct {
	UUID          string
	Name          string
	Description   string
	PriceCents    int64
	StockQuantity int64
	Category      Category
	Dimensions    Dimensions
	Manufacturer  Manufacturer
	Tags          []string
	Metadata      map[string]interface{}
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Category string

const (
	CategoryUnspecified Category = "UNSPECIFIED"
	CategoryEngine      Category = "ENGINE"
	CategoryFuel        Category = "FUEL"
	CategoryPorthole    Category = "PORTHOLE"
	CategoryWing        Category = "WING"
)

type Dimensions struct {
	Length float64 // метры
	Width  float64 // метры
	Height float64 // метры
	Weight float64 // килограммы
}

type Manufacturer struct {
	Name    string
	Country string
	Website string
}

type PartsFilter struct {
	Uuids                 []string
	Names                 []string
	Categories            []Category
	ManufacturerCountries []string
	Tags                  []string
}
