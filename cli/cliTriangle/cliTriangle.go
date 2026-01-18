package cliTriangle

import "fmt"

func InputTriangleParam() (a, b, c float64) {
	for {
		fmt.Println("Введите параметры треугольника")
		_, err := fmt.Scanln(&a, &b, &c)
		if err != nil {
			fmt.Println("Ошибка", err)
			fmt.Println("Повторите ввод")
			continue
		}
		if a <= 0 || b <= 0 || c <= 0 {
			fmt.Println("Длины сторон треугольника должны иметь положительные значения")
			continue
		}
		firstSum := a + b
		secondSum := b + c
		thirdSum := a + c

		if firstSum <= c || secondSum <= a || thirdSum <= b {
			fmt.Println("Нарушение неравенства треугольника. Повторите ввод")
			continue
		}

		break
	}
	return a, b, c
}
