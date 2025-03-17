package pkg

import (
	"context"
	"fmt"
	"fsa/internal/dto"
	"fsa/pkg/util"
	"net/http"
	"time"
)

type BargainSearcherConfig struct {
	BaseURL           string
	SearchPath        string
	SearchAPICallPath string
	AuthToken         string
}

type bargainSearch struct {
	baseURL           string
	SearchPath        string
	searchAPICallPath string
	authToken         string
}

var _ FlightSearcher = (*bargainSearch)(nil)

func NewBargainSearch(conf BargainSearcherConfig) FlightSearcher {
	bs := &bargainSearch{
		baseURL:    conf.BaseURL,
		SearchPath: conf.SearchPath,
		authToken:  conf.AuthToken,
	}
	bs.searchAPICallPath = fmt.Sprintf("%s%s", bs.baseURL, conf.SearchAPICallPath)
	return bs
}

func (bs *bargainSearch) SearchFlight(
	ctx context.Context,
	req dto.SearchTravelDTO,
) (any, error) {
	// prepare response objects
	response := dto.GroupedItineraryResponse{}
	errorResponse := make(map[string]any, 0)

	//prepare request from dto
	originDestination := make([]dto.OriginDestinationInformation, 0, len(req.Legs))
	for _, travel := range req.Legs {
		originDestination = append(originDestination, dto.OriginDestinationInformation{
			DepartureDateTime: travel.DepartureDate.Format(time.RFC3339),
			OriginLocation: dto.Location{
				LocationCode: string(travel.Origin),
			},
			DestinationLocation: dto.Location{
				LocationCode: string(travel.Destination),
			},
		})
	}

	bargainRequest := dto.BargainFinderMaxRequest{
		OTAAirLowFareSearchRQ: dto.OTAAirLowFareSearchRQ{
			OriginDestinationInformation: originDestination,
			TravelerInfoSummary:          req.TravelerSummary,
		},
	}
	err := util.MakeRequest(
		ctx,
		http.MethodPost,
		bs.searchAPICallPath,
		func(r *http.Request) {
			// set authorization token
			r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bs.authToken))
		},
		bargainRequest,
		&response,
		&errorResponse,
	)
	if err != nil {
		return response, err
	}
	return response, nil
}
