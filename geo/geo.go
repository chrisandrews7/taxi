package geo

import (
	"github.com/golang/geo/s1"
	"github.com/golang/geo/s2"
)

const (
	searchMin         = 10
	searchMax         = 16
	searchIntervalMod = 2

	searchMaxCells = 10
	// The Earth's mean radius in kilometers (according to NASA).
	earthRadiusKm = 6371.01
)

func kmToAngle(km float64) s1.Angle {
	return s1.Angle(km / earthRadiusKm)
}

// TokensFromCell returns parent cells for a given cell based on a predefined min and max level
// It ensures we use the same cells when trying to match if multiple points are in the same area
func TokensFromCell(cellID s2.CellID) []string {
	var tokens []string

	for i := searchMin; i < searchMax; i++ {
		if (searchMin % searchIntervalMod) != (i % searchIntervalMod) {
			continue
		}

		if cellID.Level() >= i {
			tokens = append(tokens, cellID.Parent(i).ToToken())
		}
	}

	return tokens
}

// TokensFromCellWithRadius returns parent cells for all neighbouring cells in its provided search radius
func TokensFromCellWithRadius(cellID s2.CellID, searchRadius float64) [][]string {
	cellIDs := CellIDsFromSearchRadius(cellID, searchRadius)

	tokens := make([][]string, 0)

	for _, cellID := range cellIDs {
		tokens = append(tokens, TokensFromCell(cellID))
	}

	return tokens
}

// CellIDsFromSearchRadius returns all the cellIDs for a given sphere around a CellID
func CellIDsFromSearchRadius(cellID s2.CellID, radiusKm float64) []s2.CellID {
	s2cap := s2.CapFromCenterAngle(cellID.Point(), kmToAngle(radiusKm))
	rc := &s2.RegionCoverer{MinLevel: 12, MaxLevel: 16, MaxCells: searchMaxCells}
	covering := rc.Covering(s2.Region(s2cap))

	var cellIDs []s2.CellID
	for _, cellID := range covering {
		cellIDs = append(cellIDs, cellID)
	}

	return cellIDs
}
