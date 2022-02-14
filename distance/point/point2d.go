package point

import "encoding/json"

type Point2d struct {
	x float64
	y float64
}

func (l Point2d) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(
		struct {
			X         float64 `json:"x"`
			Y         float64 `json:"y"`
			CreatedAt string  `json:"createdAt,omitempty"`
		}{l.x, l.y, ""})

	if err != nil {
		return nil, err
	}

	return j, err
}

func (p Point2d) X() float64 {
	return p.x
}

func (p Point2d) Y() float64 {
	return p.y
}

func NewPoint2d(x float64, y float64) *Point2d {
	return &Point2d{x: x, y: y}
}
