package main

import (
	"fmt"
	"lesson2/distance"
	"lesson2/distance/navigator"
	"lesson2/distance/point"
)

func main() {
	line := distance.NewLine2d(*point.NewPoint2d(1, 0), *point.NewPoint2d(2, 3))

	fmt.Println(line)

	distance2dLine := line.Distance()

	fmt.Printf("line2d %T имеет растояние %v\n", line, distance2dLine)

	line3d1 := distance.NewLine3d(*point.NewPoint3d(1, 0, 1), *point.NewPoint3d(0, 1, 0))

	fmt.Println(line3d1)

	line3d1Distance, _ := line3d1.Distance()
	fmt.Printf("line3d имеет растояние %v\n", line3d1Distance)

	path3d1 := distance.NewPath3d(make([]point.Point3d, 0))
	path3d1.AddPoint(*point.NewPoint3d(1, 1, 1)).AddPoint(*point.NewPoint3d(2, 3, 3)).AddPoint(*point.NewPoint3d(2, 3, 4))

	distancePath3d, _ := path3d1.Distance()
	fmt.Printf("path3d полный путь %v\n", distancePath3d)

	navi := navigator.NewNavigator([]navigator.AllDistance{path3d1, line3d1})

	fullPath, err := navi.Path()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Полный путь по навигатору %v\n", fullPath)

	PrintfLn("Path %v", fullPath)

}

func PrintfLn(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format+"\n", a...)
}
