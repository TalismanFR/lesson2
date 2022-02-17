package navigator

import (
	"errors"
	"lesson2/distance/navigator/infra"
	"lesson2/distance/point"
)

type AllDistance interface {
	Distance() (float64, error)
}

type Navigator struct {
	//distances []AllDistance
	geocoder infra.Geocoding
}

type PathInfo struct {
	placeStart  infra.GeocodeData
	placeFinish infra.GeocodeData
}

func (p PathInfo) PlaceStart() infra.GeocodeData {
	return p.placeStart
}

func (p PathInfo) PlaceFinish() infra.GeocodeData {
	return p.placeFinish
}

//func NewNavigator(distances []AllDistance) *Navigator {
//	return &Navigator{distances: distances}
//}

func NewNavigator(geocoding infra.Geocoding) *Navigator {
	return &Navigator{geocoder: geocoding}
}

//func (n Navigator) Path() (path float64, err error) {
//	for _, dist := range n.distances {
//		pathLocal, err1 := dist.Distance()
//
//		if err1 != nil {
//			return path, err1
//		}
//		path += pathLocal
//	}
//
//	return
//
//}

func (n Navigator) PathInfo(point1 point.Point2d, point2 point.Point2d) (info PathInfo, err error) {
	data1, err := n.geocoder.ReverseGeocode(point1)
	if err != nil {
		err = errors.New("Error geocode first point")
	}

	data2, err := n.geocoder.ReverseGeocode(point2)
	if err != nil {
		err = errors.New("Error geocode second point")
	}

	info = PathInfo{placeStart: data1, placeFinish: data2}
	return
}
