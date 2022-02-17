package distance

import (
	"github.com/stretchr/testify/assert"
	"lesson2/distance/point"
	"testing"
)

func TestLine3d_Distance(t *testing.T) {
	point1 := point.NewPoint3d(0, 0, 0)
	poin2 := point.NewPoint3d(0, 0, 0)
	line := NewLine3d(*point1, *poin2)
	dist, err := line.Distance()

	if err != nil {
		t.Fatalf("error calc distance %v with error %v", dist, err.Error())
	}
	if dist != 0 {
		t.Errorf("error dist value from 0 to 0 %v", dist)
	}

	type parameters struct {
		point1 point.Point3d
		point2 point.Point3d
	}
	type testCase struct {
		name   string
		params parameters
		want   float64
	}

	tests := []testCase{
		{
			"zero distance", parameters{*point.NewPoint3d(0, 0, 0), *point.NewPoint3d(0, 0, 0)}, 0,
		},
		{
			"one distance", parameters{*point.NewPoint3d(0, 0, 0), *point.NewPoint3d(1, 0, 0)}, 1,
		},
		{
			"any distance", parameters{*point.NewPoint3d(-4, 0, 2), *point.NewPoint3d(1, 3, 4)}, 6.164414002968976,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			line := NewLine3d(tc.params.point1, tc.params.point2)
			dist, err := line.Distance()
			assert.Nil(t, err)
			assert.Equal(t, dist, tc.want)
		})
	}
}

func TestNewLine3d(t *testing.T) {
	line := NewLine3d(*point.NewPoint3d(0, 0, 0), *point.NewPoint3d(1, 1, 1))
	assert.Equal(t, line.point1, *point.NewPoint3d(0, 0, 0), "point 1 should bu equal 0,0,0")
	assert.NotEqual(t, line.point1, line.point2)
}
