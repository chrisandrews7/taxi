package matcher

import (
	"github.com/chrisandrews7/taxi/geo"
	"github.com/golang/geo/s2"
)

type LatLng struct {
	Lat, Lng float64
}

func Matches(customer LatLng, driver LatLng, radius float64) []string {
	c := s2.CellIDFromLatLng(s2.LatLngFromDegrees(customer.Lat, customer.Lng))
	d := s2.CellIDFromLatLng(s2.LatLngFromDegrees(driver.Lat, driver.Lng))

	dTokens := geo.TokensFromCellWithRadius(d, radius)
	cTokens := geo.TokensFromCell(c)

	// fmt.Printf("Customer Tokens %s \n", cTokens)
	// fmt.Printf("Driver Tokens %s \n", dTokens)

	matches := make([]string, 0)
	for _, t := range dTokens {
		for _, t2 := range cTokens {
			if t[0] == t2 {
				matches = append(matches, t[0])
			}
		}
	}

	return matches
}
