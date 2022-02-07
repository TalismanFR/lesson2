package main

import (
	"bufio"
	"fmt"
	"math/cmplx"
	"os"
	"strconv"
	"strings"
)

func main() {
	var A, B, C float64

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Произошла паника ", r)
		}

	}()

	var err error
	//A, B, C, err = floatFromFile("abc")
	A, B, C, err = floatesFromFile("abc")
	//A, err = getFloatParam("A")
	if err != nil {
		fmt.Printf("Error parse float " + err.Error())
		return
	}

	//else {
	//	fmt.Println("Value A=", A)
	//}
	//
	//B, err = getFloatParam("B")
	//if err != nil {
	//	fmt.Printf("Error parse float " + err.Error())
	//	return
	//}
	//C, err = getFloatParam("C")
	//if err != nil {
	//	fmt.Printf("Error parse float " + err.Error())
	//	return
	//}

	fmt.Printf("Уравнение %.2fx^2+%.2fx+%.2f\n", A, B, C)

	a := complex(A, 0)
	b := complex(B, 0)
	c := complex(C, 0)

	sqrD := cmplx.Sqrt(b*b - 4*a*c)
	x1 := (-b + sqrD) / (2 * a)
	x2 := (-b - sqrD) / (2 * a)

	fmt.Printf("d=", sqrD)
	fmt.Printf("x1=", x1, "x2=", x2)

}

func floatesFromFile(filename string) (a, b, c float64, err error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Файл " + filename + " успешно получен")
	}

	defer fmt.Printf("первый девер\n")
	defer fmt.Printf("второй девер\n")
	defer func() {
		fmt.Printf("Закрытие файла\n")
		file.Close()
	}()

	scanner := bufio.NewScanner(file)

	var line string
	for scanner.Scan() {
		line += scanner.Text() + "\n"
	}

	fmt.Printf("%T %v\n", line, line)

	for i, code := range line {
		fmt.Printf("string step %v val=%c \n", i, code)
	}

	lines := strings.Split(line, "\n")

	fmt.Printf("%T %v\n", lines, lines)

	for i, code := range lines {
		fmt.Printf("step %v val=%v\n", i, code)

		switch i {
		case 0:
			a, _ = strconv.ParseFloat(code, 64)
		case 1:
			b, _ = strconv.ParseFloat(code, 64)
		case 2:
			c, _ = strconv.ParseFloat(code, 64)
		default:
			fmt.Printf("index ", i, "val=", code)
		}
	}

	return
}

func floatFromFile(filename string) (a, b, c float64, err error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Файл " + filename + " успешно получен")
	}

	defer fmt.Printf("первый девер\n")
	defer fmt.Printf("второй девер\n")
	defer func() {
		fmt.Printf("Закрытие файла\n")
		file.Close()
	}()

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		a, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("error parse value" + scanner.Text())
			panic("error parse")
			//return
		}
	}

	if scanner.Scan() {
		b, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("error parse value" + scanner.Text())
			return
		}
	}

	if scanner.Scan() {
		c, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("error parse value" + scanner.Text())
			return
		}
	}

	return

}

func getFloatParam(label string) (param float64, err error) {
	var val string
	fmt.Println(label + "=")
	fmt.Scanln(&param)
	param, err = strconv.ParseFloat(val, 64)
	return
}
