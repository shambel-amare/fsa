package cmd

import (
	"fsa/pkg"

	"github.com/spf13/viper"
)

func InitiatePackageLayer() pkg.FlightSearcher {
	// initiate package layer
	conf := pkg.BargainSearcherConfig{
		BaseURL:    viper.GetString("bargain.baseUrl"),
		AuthToken:  viper.GetString("bargain.authToken"),
		SearchPath: viper.GetString("bargain.searchPath"),
	}
	return pkg.NewBargainSearch(conf)
}
