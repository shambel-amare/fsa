package dto

type GroupedItineraryResponse struct {
	Version                string                  `json:"version"`
	Messages               []Message               `json:"messages"`
	Statistics             Statistics              `json:"statistics"`
	ScheduleDescs          []ScheduleDesc          `json:"scheduleDescs"`
	TaxDescs               []TaxDesc               `json:"taxDescs"`
	TaxSummaryDescs        []TaxSummaryDesc        `json:"taxSummaryDescs"`
	ObFeeDescs             []ObFeeDesc             `json:"obFeeDescs"`
	FareComponentDescs     []FareComponentDesc     `json:"fareComponentDescs"`
	ValidatingCarrierDescs []ValidatingCarrierDesc `json:"validatingCarrierDescs"`
	BaggageAllowanceDescs  []BaggageAllowanceDesc  `json:"baggageAllowanceDescs"`
	LegDescs               []LegDesc               `json:"legDescs"`
	ItineraryGroups        []ItineraryGroup        `json:"itineraryGroups"`
}

type Message struct {
	Severity string `json:"severity"`
	Type     string `json:"type"`
	Code     string `json:"code"`
	Text     string `json:"text"`
}

type Statistics struct {
	ItineraryCount int `json:"itineraryCount"`
}

type ScheduleDesc struct {
	ID              int      `json:"id"`
	Frequency       string   `json:"frequency"`
	StopCount       int      `json:"stopCount"`
	ETicketable     bool     `json:"eTicketable"`
	TotalMilesFlown int      `json:"totalMilesFlown"`
	ElapsedTime     int      `json:"elapsedTime"`
	Departure       Location `json:"departure"`
	Arrival         Location `json:"arrival"`
	Carrier         Carrier  `json:"carrier"`
}

type Carrier struct {
	Marketing             string    `json:"marketing"`
	MarketingFlightNumber int       `json:"marketingFlightNumber"`
	Operating             string    `json:"operating"`
	OperatingFlightNumber int       `json:"operatingFlightNumber"`
	Equipment             Equipment `json:"equipment"`
}

type Equipment struct {
	Code            string `json:"code"`
	TypeForFirstLeg string `json:"typeForFirstLeg"`
	TypeForLastLeg  string `json:"typeForLastLeg"`
}

type TaxDesc struct {
	ID                int     `json:"id"`
	Code              string  `json:"code"`
	Amount            float64 `json:"amount"`
	Currency          string  `json:"currency"`
	Description       string  `json:"description"`
	PublishedAmount   float64 `json:"publishedAmount"`
	PublishedCurrency string  `json:"publishedCurrency"`
	Station           string  `json:"station"`
	Country           string  `json:"country,omitempty"`
}

type TaxSummaryDesc struct {
	ID                int     `json:"id"`
	Code              string  `json:"code"`
	Amount            float64 `json:"amount"`
	Currency          string  `json:"currency"`
	Description       string  `json:"description"`
	PublishedAmount   float64 `json:"publishedAmount"`
	PublishedCurrency string  `json:"publishedCurrency"`
	Station           string  `json:"station"`
	Country           string  `json:"country,omitempty"`
}

type ObFeeDesc struct {
	ID       int     `json:"id"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type FareComponentDesc struct {
	ID                          int              `json:"id"`
	GoverningCarrier            string           `json:"governingCarrier"`
	FareAmount                  float64          `json:"fareAmount"`
	FareCurrency                string           `json:"fareCurrency"`
	FareBasisCode               string           `json:"fareBasisCode"`
	FarePassengerType           string           `json:"farePassengerType"`
	PublishedFareAmount         float64          `json:"publishedFareAmount"`
	PublishedFareCurrency       string           `json:"publishedFareCurrency"`
	Directionality              string           `json:"directionality"`
	ApplicablePricingCategories string           `json:"applicablePricingCategories"`
	VendorCode                  string           `json:"vendorCode"`
	FareTypeBitmap              string           `json:"fareTypeBitmap"`
	FareType                    string           `json:"fareType"`
	FareTariff                  string           `json:"fareTariff"`
	FareRule                    string           `json:"fareRule"`
	CabinCode                   string           `json:"cabinCode"`
	Segments                    []SegmentWrapper `json:"segments"`
}

type SegmentWrapper struct {
	Segment Segment `json:"segment"`
}

type Segment struct {
	Stopover bool `json:"stopover"`
}

type ValidatingCarrierDesc struct {
	ID               int         `json:"id"`
	SettlementMethod string      `json:"settlementMethod"`
	NewVcxProcess    bool        `json:"newVcxProcess"`
	Default          CarrierCode `json:"default"`
}

type CarrierCode struct {
	Code string `json:"code"`
}

type BaggageAllowanceDesc struct {
	ID         int `json:"id"`
	PieceCount int `json:"pieceCount"`
}

type LegDesc struct {
	ID          int           `json:"id"`
	ElapsedTime int           `json:"elapsedTime"`
	Schedules   []ScheduleRef `json:"schedules"`
}

type ScheduleRef struct {
	Ref int `json:"ref"`
}

type ItineraryGroup struct {
	GroupDescription GroupDescription `json:"groupDescription"`
	Itineraries      []Itinerary      `json:"itineraries"`
}

type GroupDescription struct {
	LegDescriptions []LegDescription `json:"legDescriptions"`
}

type LegDescription struct {
	DepartureDate     string `json:"departureDate"`
	DepartureLocation string `json:"departureLocation"`
	ArrivalLocation   string `json:"arrivalLocation"`
}

type Itinerary struct {
	ID                 int                  `json:"id"`
	PricingSource      string               `json:"pricingSource"`
	Legs               []LegRef             `json:"legs"`
	PricingInformation []PricingInformation `json:"pricingInformation"`
}

type LegRef struct {
	Ref int `json:"ref"`
}

type PricingInformation struct {
	PricingSubsource string `json:"pricingSubsource"`
	Offer            Offer  `json:"offer"`
	Fare             Fare   `json:"fare"`
}

type Offer struct {
	OfferID    string `json:"offerId"`
	TimeToLive int    `json:"timeToLive"`
	Source     string `json:"source"`
}

type Fare struct {
	OfferItemID           string                 `json:"offerItemId"`
	MandatoryInd          bool                   `json:"mandatoryInd"`
	ServiceID             string                 `json:"serviceId"`
	ValidatingCarrierCode string                 `json:"validatingCarrierCode"`
	Vita                  bool                   `json:"vita"`
	ETicketable           bool                   `json:"eTicketable"`
	LastTicketDate        string                 `json:"lastTicketDate"`
	LastTicketTime        string                 `json:"lastTicketTime"`
	GoverningCarriers     string                 `json:"governingCarriers"`
	PassengerInfoList     []PassengerInfoWrapper `json:"passengerInfoList"`
	TotalFare             TotalFare              `json:"totalFare"`
	ValidatingCarriers    []ValidatingCarrierRef `json:"validatingCarriers"`
}

type PassengerInfoWrapper struct {
	PassengerInfo PassengerInfo `json:"passengerInfo"`
}

type PassengerInfo struct {
	PassengerType      string               `json:"passengerType"`
	PassengerNumber    int                  `json:"passengerNumber"`
	NonRefundable      bool                 `json:"nonRefundable"`
	FareComponents     []FareComponentRef   `json:"fareComponents"`
	Taxes              []TaxRef             `json:"taxes"`
	TaxSummaries       []TaxSummaryRef      `json:"taxSummaries"`
	ObFees             []ObFeeRef           `json:"obFees"`
	CurrencyConversion CurrencyConversion   `json:"currencyConversion"`
	PassengerTotalFare PassengerTotalFare   `json:"passengerTotalFare"`
	BaggageInformation []BaggageInformation `json:"baggageInformation"`
}

type FareComponentRef struct {
	Ref int `json:"ref"`
}

type TaxRef struct {
	Ref int `json:"ref"`
}

type TaxSummaryRef struct {
	Ref int `json:"ref"`
}

type ObFeeRef struct {
	Ref int `json:"ref"`
}

type CurrencyConversion struct {
	From             string  `json:"from"`
	To               string  `json:"to"`
	ExchangeRateUsed float64 `json:"exchangeRateUsed"`
}

type PassengerTotalFare struct {
	TotalFare            float64 `json:"totalFare"`
	TotalTaxAmount       float64 `json:"totalTaxAmount"`
	Currency             string  `json:"currency"`
	BaseFareAmount       float64 `json:"baseFareAmount"`
	BaseFareCurrency     string  `json:"baseFareCurrency"`
	EquivalentAmount     float64 `json:"equivalentAmount"`
	EquivalentCurrency   string  `json:"equivalentCurrency"`
	ConstructionAmount   float64 `json:"constructionAmount"`
	ConstructionCurrency string  `json:"constructionCurrency"`
	ExchangeRateOne      float64 `json:"exchangeRateOne"`
}

type BaggageInformation struct {
	ProvisionType string       `json:"provisionType"`
	AirlineCode   string       `json:"airlineCode"`
	Segments      []SegmentID  `json:"segments"`
	Allowance     AllowanceRef `json:"allowance"`
}

type SegmentID struct {
	ID int `json:"id"`
}

type AllowanceRef struct {
	Ref int `json:"ref"`
}

type ValidatingCarrierRef struct {
	Ref int `json:"ref"`
}

type TotalFare struct {
	TotalPrice           float64 `json:"totalPrice"`
	TotalTaxAmount       float64 `json:"totalTaxAmount"`
	Currency             string  `json:"currency"`
	BaseFareAmount       float64 `json:"baseFareAmount"`
	BaseFareCurrency     string  `json:"baseFareCurrency"`
	ConstructionAmount   float64 `json:"constructionAmount"`
	ConstructionCurrency string  `json:"constructionCurrency"`
	EquivalentAmount     float64 `json:"equivalentAmount"`
	EquivalentCurrency   string  `json:"equivalentCurrency"`
}
