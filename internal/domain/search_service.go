package domain

import (
	"context"
	"fsa/internal/dto"
	"fsa/pkg"
)

type searchServiceUseCase struct {
	searcher pkg.FlightSearcher
}

func NewSearchService(sr pkg.FlightSearcher) SearchService {
	return &searchServiceUseCase{
		searcher: sr,
	}
}

func (s *searchServiceUseCase) SearchFlight(
	ctx context.Context,
	req dto.SearchTravelDTO,
) (any, error) {
	// make the search call with the searcher
	// create the bargain search request from the the request data

	responseData, err := s.searcher.SearchFlight(ctx, req)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}
