package main

import (
	"Figures/cli"
	"Figures/cli/cliCircle"
	"Figures/cli/cliRectangle"
	"Figures/cli/cliTriangle"
	"Figures/figure"
	"Figures/service"
	"fmt"
)

func main() {
	var choice int
	flag := true

	for flag {
		cli.CallMainMenu()
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			var a, b = cliRectangle.InputRectangleParam()
			fmt.Println("Расчеты: ")
			f := figure.NewRectangle(a, b)
			resSquare, err := f.Square()
			service.WrapError("Ошибка расчета площади", err)
			resPerimeter, err := f.Perimeter()
			service.WrapError("Ошибка расчета периметра", err)
			cli.PrintResult(resSquare, resPerimeter)
		case 2:
			var a, b, c = cliTriangle.InputTriangleParam()
			fmt.Println("Расчеты: ")
			f := figure.NewTriangle(a, b, c)
			resSquare, err := f.Square()
			service.WrapError("Ошибка расчета площади", err)
			resPerimeter, err := f.Perimeter()
			service.WrapError("Ошибка расчета периметра", err)
			cli.PrintResult(resSquare, resPerimeter)
		case 3:
			fmt.Println("Введите параметры окружности")
			r := cliCircle.InputCircleParam()
			fmt.Println("Расчеты: ")
			f := figure.NewCircle(r)
			resSquare, err := f.Square()
			service.WrapError("Ошибка расчета площади", err)
			resPerimeter, err := f.Perimeter()
			service.WrapError("Ошибка расчета периметра", err)
			cli.PrintResult(resSquare, resPerimeter)
		case 4:
			fmt.Println("Завершение работы программы")
			flag = false
		default:
			fmt.Println("Введен неверный номер фигуры, начните заново")
		}
		flag = cli.JobContinue()
	}

}
