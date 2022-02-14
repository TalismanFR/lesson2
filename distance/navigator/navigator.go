package navigator

type AllDistance interface {
	Distance() (float64, error)
}

type Navigator struct {
	distances []AllDistance
}

func NewNavigator(distances []AllDistance) *Navigator {
	return &Navigator{distances: distances}
}

func (n Navigator) Path() (path float64, err error) {
	for _, dist := range n.distances {
		pathLocal, err1 := dist.Distance()

		if err1 != nil {
			return path, err1
		}
		path += pathLocal
	}

	return

}
