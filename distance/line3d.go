package distance

import (
	"lesson2/distance/point"
	"math"
)

type Line3d struct {
	point1 point.Point3d
	point2 point.Point3d
}

func (l Line3d) Point1() point.Point3d {
	return l.point1
}

func (l Line3d) Point2() point.Point3d {
	return l.point2
}

func NewLine3d(point1 point.Point3d, point2 point.Point3d) *Line3d {
	return &Line3d{point1: point1, point2: point2}
}

func (l Line3d) Distance() (distance float64, err error) {
	distance = math.Sqrt(math.Pow(l.point2.X()-l.point1.X(), 2) + math.Pow(l.point2.Y()-l.point1.Y(), 2) + math.Pow(l.point2.Z()-l.point1.Z(), 2))
	return
}
