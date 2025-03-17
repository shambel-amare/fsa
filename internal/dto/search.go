package dto

import (
	"time"
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
type OriginDestinationInfo struct {
	TravelType    TravelType
	Origin        IATA_CODE
	Destination   IATA_CODE
	DepartureDate time.Time
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
