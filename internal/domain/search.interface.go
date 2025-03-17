package domain

import (
	"context"
	"fsa/internal/dto"
)

type SearchService interface {
	SearchFlight(ctx context.Context, req dto.SearchTravelDTO) (any, error)
}
