/* Работа функций в Go */
package main

import (
	"errors"
	"fmt"
)

const pi = 3.1415 // тип float32

func main() {
	fmt.Printf("Привет Go!!!\n")
	printCircleArea(-2)
	printCircleArea(4)

}

func printCircleArea(radius int) {
	circleArea, err := calculateCircleArea(radius)
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("Радиус круга положительный, работаем!")
	}

	fmt.Printf("Радиус круга: %d см. \n", radius)
	fmt.Printf("Формула для расчета площади круга: A=pi*r^2\n")
	fmt.Printf("Площадь круга: %f32 см. кв.\n\n", circleArea)

}

func calculateCircleArea(radius int) (float32, error) {
	if radius <= 0 {
		return float32(0), errors.New("Радиус круга не может быть отрицательным")
	}
	return float32(radius) * float32(radius) * pi, nil
}
