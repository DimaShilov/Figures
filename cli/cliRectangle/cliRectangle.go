package cliRectangle

import "fmt"

func InputRectangleParam() (a, b float64) {
	for {
		fmt.Println("Введите параметры прямоугольника")
		_, err := fmt.Scanln(&a, &b)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Введено недостаточно параметров. Повторите ввод.")
			continue
		}
		if a <= 0 || b <= 0 {
			fmt.Println("Длины сторон прямоугольника должны иметь положительные значения")
			continue
		}
		break
	}
	return a, b
}
