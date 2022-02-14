package point

type Point3d struct {
	x float64
	y float64
	z float64
}

func (p Point3d) X() float64 {
	return p.x
}

func (p Point3d) Y() float64 {
	return p.y
}

func (p Point3d) Z() float64 {
	return p.z
}

func NewPoint3d(x float64, y float64, z float64) *Point3d {
	return &Point3d{x: x, y: y, z: z}
}
