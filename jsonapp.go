package main

import (
	"encoding/json"
	"fmt"
	"lesson2/distance"
	"lesson2/distance/point"
)

func main() {
	line := distance.NewLine2d(*point.NewPoint2d(1, 0), *point.NewPoint2d(2, 3))

	b, _ := json.Marshal(line)

	fmt.Println(string(b))
}
