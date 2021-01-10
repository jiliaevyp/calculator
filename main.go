package main

import (
	"fmt"
	"strconv"
)

var arg1, arg2, result float64
var mach string
var errcalcul int // 0-ошибок нет, 1-ошибка в числе, 2-ошибка в операторе, 3-деление на ноль

const zeroerr = 0
const errnumbe = 1
const erroper = 2
const errdivzero = 3
const plus string = "+"
const minus string = "-"
const mult string = "*"
const div string = "/"

func callcul() (float64, int) {
	var err int
	arg1, err = inparg(1)
	if err > 0 {
		return 0, errnumbe
	}
	mach, err = inpmachen()
	if err > 0 {
		return 0, erroper
	}
	arg2, err = inparg(2)
	if err > 0 {
		return 0, errnumbe
	}
	result, err = operation()
	return result, err
}

func inparg(num int) (float64, int) {
	var res float64
	var intString string
	var err int
	err = 0
	if num == 1 {
		fmt.Println("Введите первое число")
	} else {
		fmt.Println("Введите второе число")
	}
	fmt.Scanf(
		"%s\n",
		&intString,
	)
	res, errconv := strconv.ParseFloat(intString, 64)
	if errconv != nil {
		err = errnumbe
		res = 0
	}
	return res, err
}

func inpmachen() (string, int) {
	var err int
	err = 1
	for err > 0 {
		fmt.Println("-----------------------------")
		fmt.Println("Выберите действие:")
		fmt.Println(plus, " сложить")
		fmt.Println(minus, " вычесть")
		fmt.Println(mult, " умножить")
		fmt.Println(div, " разделить")
		fmt.Println("-----------------------------")
		fmt.Scanf(
			"%s/n",
			&mach,
		)
		if mach == plus || mach == minus || mach == mult || mach == div {
			err = 0
			switch mach {
			case plus:
				fmt.Println(" сложить")
			case minus:
				fmt.Println(" вычесть")
			case mult:
				fmt.Println(" умножить")
			case div:
				fmt.Println(" разделить")
			}
		} else {
			fmt.Println("*** Введите правильный символ!!! ***")
		}
	}
	fmt.Println("-----------------------------")
	return mach, err
}

func operation() (float64, int) {
	var res float64
	var err int
	res = 0
	err = 0
	switch mach {
	case plus:
		res = arg1 + arg2
	case minus:
		res = arg1 - arg2
	case mult:
		res = arg1 * arg2
	case div:
		{
			if arg2 == 0 && mach == div {
				err = errdivzero // деление невозможно
			} else {
				res = arg1 / arg2 // деление возможно
			}
		}
	}
	return res, err // выход
}

func printresult(err int) {

	switch err {
	case zeroerr:
		{
			fmt.Println("-----------------------------")
			fmt.Print("Результат:", "\n")
			fmt.Printf("%.5f", arg1)
			fmt.Print(" ", mach, " ")
			fmt.Printf("%.5f", arg2)
			fmt.Print(" = ")
			fmt.Printf("%.5f", result)
			fmt.Print("\n")
		}
	case errnumbe:
		{
			fmt.Println("\n", "*** К сожалению, вы ввели не число!!")
			fmt.Println("*** Операция прекращена ***")
		}
	case erroper:
		{
			fmt.Println("\n", "*** К сожалению, вы ввели неверный символ ***")
			fmt.Println("*** Операция прекращена ***")
		}
	case errdivzero:
		{
			fmt.Println("-----------------------------")
			fmt.Println("*** Деление на ноль невозможно ***")
			fmt.Println("*** Операция прекращена ***")
		}
	}
	fmt.Print("-----------------------------", "\n", "\n")
}

func main() {
	var komand string
	arg1 = 0
	arg2 = 0
	result = 0
	mach = "?"
	errcalcul = 0
	komand = "Y"
	for komand == "Y" || komand == "y" || komand == "Н" || komand == "н" {
		fmt.Println("-----------------------------")
		fmt.Println("|       Калькулятор          |")
		fmt.Println("|  Считать, не пересчитать!  |")
		fmt.Println("|                            |")
		fmt.Println("|   (c) jiliaevyp@gmail.com  |")
		fmt.Println("-----------------------------")
		result, errcalcul = callcul()
		printresult(errcalcul)
		fmt.Println("Продолжить? (Y)")
		fmt.Println("Закончить? (Enter)")
		komand = ""
		fmt.Scanf(
			"%s\n",
			&komand,
		)

	}
	fmt.Println("Рад был с Вами пработать!")
	fmt.Print("Обращайтесь в любое время без колебаний!", "\n", "\n")
}
