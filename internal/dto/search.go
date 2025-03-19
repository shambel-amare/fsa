package dto

import (
	"fmt"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Two letter codes for countries/airports
type IATA_CODE string

const (
	ET IATA_CODE = "ET"
)

type TravelType int

const (
	OneWayTrip TravelType = iota
	RoundTrip
)

type TravlerType int

const (
	Adult TravlerType = iota
	Child
	//....others
)

type GenderType int

const (
	MALE GenderType = iota
	FEMALE
)

type SearchTravelDTO struct {
	Legs            []OriginDestinationInfo `json:"legs"`
	TravelerSummary TravelerInfoSummary     `json:"traveler_summary"`
}

// TODO: add validation
func (s SearchTravelDTO) Validate() error {
	return validation.ValidateStruct(
		&s,
		validation.Field(
			&s.TravelerSummary,
			validation.Required,
			validation.By(func(value interface{}) error {
				if travelerSummary, ok := value.(TravelerInfoSummary); ok {
					return travelerSummary.Validate() // Validate the nested TravelerInfoSummary
				}
				return fmt.Errorf("invalid TravelerInfoSummary type")
			}),
		),
		validation.Field(
			&s.Legs,
			validation.Required,
			validation.Each(validation.By(func(value interface{}) error {
				if leg, ok := value.(OriginDestinationInfo); ok {
					return leg.Validate() // Validate each item in the slice
				}
				return fmt.Errorf("invalid item type")
			})),
		),
	)
}

type OriginDestinationInfo struct {
	TravelType    TravelType
	Origin        IATA_CODE
	Destination   IATA_CODE
	DepartureDate time.Time
}

func (o OriginDestinationInfo) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.Origin, validation.Required),
		validation.Field(&o.Destination, validation.Required),
		validation.Field(&o.DepartureDate, validation.Required),
		validation.Field(&o.TravelType, validation.Required),
	)
}

type ConfirmTravelDTO struct {
	Travlers   []TravelerDTO
	TravelID   string
	TravelType TravelType
	ReturnDate time.Time
}

type TravelerDTO struct {
	Type   TravlerType
	Gender GenderType
	Name   string
	Age    int
}
