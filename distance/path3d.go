package distance

import (
	"errors"
	"fmt"
	"lesson2/distance/point"
)

type Path3d struct {
	points []point.Point3d
}

func NewPath3d(points []point.Point3d) *Path3d {
	return &Path3d{points: points}
}

func (l *Path3d) AddPoint(point point.Point3d) *Path3d {
	l.points = append(l.points, point)
	return l
}

func (l Path3d) Distance() (distance float64, err error) {

	if len(l.points) < 2 {
		return distance, errors.New("Кол-во точек меньше двух")
	}

	for i, p := range l.points {
		if i == len(l.points)-1 {
			break
		}
		distanceTwoPoint, _ := l.distanceTwoPoint(p, l.points[i+1])
		fmt.Println("Растояние между точками", p, l.points[i+1], " равно ", distanceTwoPoint)
		distance += distanceTwoPoint

	}
	return
}

func (l Path3d) distanceTwoPoint(x point.Point3d, y point.Point3d) (float64, error) {
	line := NewLine3d(x, y)
	return line.Distance()
}
