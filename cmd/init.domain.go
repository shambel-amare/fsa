package cmd

import (
	"fsa/internal/domain"
	"fsa/pkg"
)

func InitiateDomainLayer(searcher pkg.FlightSearcher) domain.SearchService {
	// Initialize the domain layer
	return domain.NewSearchService(searcher)
}
