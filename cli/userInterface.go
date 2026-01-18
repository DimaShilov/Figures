package cli

import "fmt"

func PrintResult(s, p float64) {
	fmt.Println("Площадь равна ", s)
	fmt.Println("Периметр равен ", p)
}

func CallMainMenu() {
	fmt.Println("Выберите фигуру для расчета площади и периметра")
	fmt.Print(
		"1. Прямоугольник\n",
		"2. Треугольник\n",
		"3. Круг\n",
		"4. Завершить работу программы\n",
	)
}

func JobContinue() bool {
	fmt.Println("Хотите продолжить? (Y/y). Для выхода нажмите любую другую кнопку.")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Ошибка обработки ввода. Завершение работы программы")
		return false
	}
	if input == "y" || input == "Y" {
		return true
	}
	return false
}
