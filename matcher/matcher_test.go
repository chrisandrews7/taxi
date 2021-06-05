package matcher_test

import (
	"testing"

	"github.com/chrisandrews7/taxi/matcher"
)

func TestMatches(t *testing.T) {
	var tests = []struct {
		customer        matcher.LatLng
		driver          matcher.LatLng
		radius          float64
		expectedMatches int
	}{
		{
			matcher.LatLng{40.728191, -73.993436},
			matcher.LatLng{40.783058, -73.971252},
			10.00,
			1,
		},
		{
			matcher.LatLng{51.509425, -0.139943},
			matcher.LatLng{51.490567, -0.185605},
			1.00,
			0,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := matcher.Matches(test.customer, test.driver, test.radius)
			if totalMatches := len(got); totalMatches != test.expectedMatches {
				t.Errorf("Matches() got %v, want %v", totalMatches, test.expectedMatches)
			}
		})
	}
}
