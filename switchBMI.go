package main

import "fmt"

func calculateBmi(height, weight float64) {
	bmi := weight / (height * height)
	fmt.Printf("Hello, your bmi is %.2f \n", bmi)
	switch {
	case bmi < 18.5:
		fmt.Printf("Underweight leh")
	case bmi <= 24.9:
		fmt.Printf("Healthy Weight!")
	case bmi <= 29.9:
		fmt.Printf("Overweight!!")
	case bmi <= 34.9:
		fmt.Printf("Overweight!!!")
	case bmi <= 39.9:
		fmt.Printf("Severly Obese!!!")
	default:
		fmt.Printf("Morbidly Obese sia!!!!")
	}
}

func main() {
	var height, weight float64

	fmt.Println("Give me your height (in m):")
	fmt.Scanln(&height)
	fmt.Println("Give me your weight (in kg):")
	fmt.Scanln(&weight)

	calculateBmi(height, weight)
}
