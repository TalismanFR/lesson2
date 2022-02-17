package infra

import "lesson2/distance/point"

type Geocoding interface {
	Geocode(address string) (point point.Point2d, err error)
	ReverseGeocode(point point.Point2d) (data GeocodeData, err error)
}

type GeocodeData struct {
	Point      point.Point2d
	Address    string
	PostalCode string
}
