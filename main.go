package main

import (
	"fmt"
	"strconv"
)

var arg1, arg2, result int
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
const opererr string = "?"

func callcul() (int, int) {
	var err int
	const opererr string = "?"
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

func inparg(num int) (int, int) {
	var res int
	var intString string
	var err int
	err = 0
	if num == 1 {
		fmt.Println("Введите первое число")
	} else {
		fmt.Println("Введите второе число")
	}
	fmt.Scanf("%s\n", &intString)
	res, errconv := strconv.Atoi(intString) // Десятичные значения
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
		fmt.Println("Выберите действие")
		fmt.Println(plus, " сложить,", minus, " вычесть,", mult, " умножить,", div, " разделить")
		fmt.Scanf("%s\n", &mach)
		if len(mach) == 1 && (mach == plus || mach == minus || mach == mult || mach == div) {
			err = 0
		}
	}
	return mach, err
}

func operation() (int, int) {
	var res int
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
			fmt.Print("Результат:", "\n")
			fmt.Print(arg1, " ", mach, " ", arg2, " = ", result, "\n")
		}
	case errnumbe:
		{
			fmt.Println("\n", "*** К сожалению, вы ввели не число!!")
			fmt.Println("*** Операция прекращена ***")
		}
	case erroper:
		{
			fmt.Println("\n", "*** К сожлению, вы ввели неверный символ ***")
			fmt.Println("*** Операция прекращена ***")
		}
	case errdivzero:
		{
			fmt.Println("\n", "*** Деление на ноль невозможно ***")
			fmt.Println("*** Операция прекращена ***")
		}
	}
	fmt.Print("-------------------------------", "\n")
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
		fmt.Println("-----------------------------")
		result, errcalcul = callcul()
		printresult(errcalcul)
		fmt.Println("Продолжить? (Y)")
		fmt.Println("Закончить? (Enter)")
		komand = ""
		fmt.Scanf("%s\n", &komand)

	}
	fmt.Println("Рад был с Вами пработать!")
	fmt.Println("Ждем Вас снова с нетерпением!", "\n\n")
}
