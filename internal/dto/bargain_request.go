package dto

type BargainFinderMaxRequest struct {
	OTAAirLowFareSearchRQ OTAAirLowFareSearchRQ `json:"OTA_AirLowFareSearchRQ"`
}
type OTAAirLowFareSearchRQ struct {
	AvailableFlightsOnly         bool                           `json:"AvailableFlightsOnly"`
	Version                      string                         `json:"Version"`
	ResponseVersion              string                         `json:"ResponseVersion"`
	RequestType                  string                         `json:"RequestType"` //[GIR-JSON]
	POS                          POS                            `json:"POS"`
	OriginDestinationInformation []OriginDestinationInformation `json:"OriginDestinationInformation"`
	TravelPreferences            TravelPreferences              `json:"TravelPreferences"`
	TravelerInfoSummary          TravelerInfoSummary            `json:"TravelerInfoSummary"`
	TPAExtensions                TPAExtensions                  `json:"TPA_Extensions"`
}

type POS struct {
	Source []Source `json:"Source"`
}

type Source struct {
	FixedPCC       bool        `json:"FixedPCC"`
	PseudoCityCode string      `json:"PseudoCityCode"`
	RequestorID    RequestorID `json:"RequestorID"`
}

type RequestorID struct {
	Type        string      `json:"Type"`
	ID          string      `json:"ID"`
	CompanyName CompanyName `json:"CompanyName"`
}

type CompanyName struct {
	Code string `json:"Code"`
}

type OriginDestinationInformation struct {
	DepartureDateTime   string   `json:"DepartureDateTime"`
	OriginLocation      Location `json:"OriginLocation"`
	DestinationLocation Location `json:"DestinationLocation"`
}

type Location struct {
	LocationCode string `json:"LocationCode"`
}

type TravelPreferences struct {
	MaxStopsQuantity int `json:"MaxStopsQuantity"`
	AncillaryFees    struct {
		AncillaryFeeGroup []struct {
			Code  string `json:"Code"`
			Count string `json:"Count"`
		} `json:"AncillaryFeeGroup"`
		Enabled bool `json:"Enabled"`
		Summary bool `json:"Summary"`
	}
	Baggage struct {
		CarryOnInfo       bool   `json:"CarryOnInfo"`
		Description       string `json:"Description"`
		FreeCarryOn       bool   `json:"FreeCarryOn"`
		FreePieceRequired bool   `json:"FreePieceRequired"`
		RequestType       string `json:"RequestType"`     //Specifies baggage information type. "A" - allowance only; "C" - allowance and charges; "N" - no baggage information. [ A, C, N ]
		RequestedPieces   int    `json:"RequestedPieces"` //[1, 4]
	} `json:"Baggage"`
	CabinPref []struct {
		Cabin       string `json:"Cabin"`       //[ PremiumFirst, First, PremiumBusiness, Business, PremiumEconomy, Economy, Y, S, C, J, F, P ]
		PreferLevel string `json:"PreferLevel"` //[ Preferred ]
	} `json:"CabinPref"`
	EticketDesired bool `json:"ETicketDesired"`
	FlightTypePref struct {
		MaxConnections int `json:"MaxConnections"`
	} `json:"FlightTypePref"`
	Hybrid                bool `json:"Hybrid"`
	LookForAlternative    bool `json:"LookForAlternative"`
	SpanishFamilyDiscount struct {
		Level int `json:"Level"` // [1,2]
	} `json:"SpanishFamilyDiscount"`
	TPAExtensions struct {
		AdditionalFareLimit struct {
			Value int `json:"Value"`
		}
		AvoidCabinDowngrade string `json:"AvoidCabinDowngrade"` //[ All, Main, None ]
		ClassOfService      []struct {
			Code        string `json:"Code"`
			PreferLevel string `json:"PreferLevel"` //[ Unacceptable, Preferred ]
		}
	} `json:"TPAExtensions"`
	VendorPref              []VendorPref `json:"VendorPref"`
	ValidInterlineTicket    bool         `json:"ValidInterlineTicket"`
	VendorPrefApplicability []struct{}
	VendorPrefPairing       []struct{}
}

type VendorPref struct {
	Code string `json:"Code"`
}

type TravelerInfoSummary struct {
	AirTravelerAvail  []AirTravelerAvail `json:"AirTravelerAvail"`
	TravelPreferences TravelPreferences  `json:"TravelPreferences"`
}

type AirTravelerAvail struct {
	PassengerTypeQuantity []PassengerTypeQuantity `json:"PassengerTypeQuantity"`
}

type PassengerTypeQuantity struct {
	PassengerId   string             `json:"PassengerId"`
	PersonName    PersonName         `json:"PersonName"`
	Code          string             `json:"Code"`
	Quantity      int                `json:"Quantity"`
	TPAExtensions []VoluntaryChanges `json:"TPA_Extensions"`
}
type VoluntaryChanges struct {
	Match    string `json:"Currency"` //[ All, Any, Info ]
	Penality []struct {
		Amount       int    `json:"Amount"`
		Application  string `json:"Application"` //[ After, Before ]
		CurrencyCode string `json:"CurrencyCode"`
		Exclude      bool   `json:"Exclude"`
		Type         string `json:"Type"` // [Refund,Exchange]
	} `json:"Penality"`
}
type PersonName struct {
	GivenName string `json:"GivenName"`
	Surname   string `json:"Surname"`
}
type TPAExtensions struct {
	IntelliSellTransaction IntelliSellTransaction `json:"IntelliSellTransaction"`
}

type IntelliSellTransaction struct {
	RequestType RequestType `json:"RequestType"`
}

type RequestType struct {
	Name string `json:"Name"`
}
