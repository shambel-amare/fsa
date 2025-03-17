package pkg

import (
	"context"
	"fsa/internal/dto"
)

type FlightSearcher interface {
	SearchFlight(
		ctx context.Context,
		req dto.SearchTravelDTO,
	) (any, error)
}
