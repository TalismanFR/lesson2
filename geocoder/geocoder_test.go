package geocoder

import (
	"lesson2/distance/point"
	"testing"
)

func TestGeocoder_ReverseGeocode(t *testing.T) {
	geocoder := NewGeocoder()

	geocoder.ReverseGeocode(*point.NewPoint2d(55.878, 37.653))
}
