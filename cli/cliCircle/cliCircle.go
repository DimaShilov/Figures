package cliCircle

import "fmt"

func InputCircleParam() (r float64) {
	for {
		_, err := fmt.Scan(&r)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Ошибка чтения. Повторите ввод.")
			continue
		}
		if r <= 0 {
			fmt.Println("Длина радиуса должна быть положительной. Повторите ввод")
			continue
		}
		break
	}
	return r
}
