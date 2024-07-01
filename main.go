package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var roman_nums = map[int]string{
	100: "C",
	90:  "XC",
	50:  "L",
	40:  "XL",
	10:  "X",
	9:   "IX",
	8:   "VIII",
	7:   "VII",
	6:   "VI",
	5:   "V",
	4:   "IV",
	3:   "III",
	2:   "II",
	1:   "I",
}
var nums_in_right_order = [14]int{100, 90, 50, 40, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
var ten_nums = [10]string{roman_nums[1], roman_nums[2], roman_nums[3], roman_nums[4], roman_nums[5], roman_nums[6], roman_nums[7], roman_nums[8], roman_nums[9], roman_nums[10]}

const (
	WELCOME_TXT         = "Пж, введите строку для расчета (тоьлко арабские целые или римские числа из интервала 1..10 и арифметические операции между ними +, -, * или /): "
	RESULT_TXT          = "Результат расчета: "
	ERR_MIXED_NUMS      = "Калькулятор умеет работать только с арабскими или римскими цифрами одновременно!"
	ERR_OUT_OF_RANGE    = "Введены не подходящие числа!"
	ERR_INCORRECT_ENTER = "Строка для расчета - не корректный ввод!"
	ERR_ROMAN_NUMS      = "Результатом работы калькулятора с римскими числами могут быть только положительные числа, т.к. римских чисел меньше 1 не бывает!"
)

func int2roman(x int) string {
	var res string
	for _, i := range nums_in_right_order {
		s := roman_nums[i]
		z := x / i
		if z > 0 && i > 9 {
			x -= z * i
			for j := 1; j <= z; j++ {
				res += s
			}
		} else if x == i && i < 10 {
			x -= i
			res += s
		}
	}
	return res
}

func roman2int(x string) (int, error) {
	for i := 0; i < len(ten_nums); i++ {
		if ten_nums[i] == x {
			return i + 1, nil
		}
	}

	return 0, errors.New(ERR_OUT_OF_RANGE)
}

func parse_str(str string) (string, []string, error) {

	var res []string

	str = clean_it_up(str)

	// +
	res = strings.Split(str, "+")
	if len(res) == 2 {
		res[0] = clean_it_up(res[0])
		res[1] = clean_it_up(res[1])
		return "+", res, nil
	}

	// -
	res = strings.Split(str, "-")
	if len(res) == 2 {
		res[0] = clean_it_up(res[0])
		res[1] = clean_it_up(res[1])
		return "-", res, nil
	}

	// *
	res = strings.Split(str, "*")
	if len(res) == 2 {
		res[0] = clean_it_up(res[0])
		res[1] = clean_it_up(res[1])
		return "*", res, nil
	}

	// /
	res = strings.Split(str, "/")
	if len(res) == 2 {
		res[0] = clean_it_up(res[0])
		res[1] = clean_it_up(res[1])
		return "/", res, nil
	}

	return "", res, errors.New(ERR_INCORRECT_ENTER)
}

func clean_it_up(str string) string {
	str = strings.TrimSpace(str)
	str = strings.ToUpper(str)
	//str = strings.ReplaceAll(str, " ", "")
	return str
}

func math_action(a int, b int, operator string, roman bool) int {
	var res int

	if a > 10 || b > 10 || a < 1 || b < 1 {
		panic(ERR_OUT_OF_RANGE)
	}

	switch operator {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	}

	if roman && res < 1 {
		panic(ERR_ROMAN_NUMS)
	}

	return res
}

func main() {
	input_reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(WELCOME_TXT)
		txt, _ := input_reader.ReadString('\n')
		txt = clean_it_up(txt)
		operator, strs, e := parse_str(txt)
		if e != nil {
			panic(ERR_INCORRECT_ENTER)
		}
		res := "NIL"
		a, ea := strconv.Atoi(strs[0])
		b, eb := strconv.Atoi(strs[1])
		if ea != nil || eb != nil {
			a, ea := roman2int(strs[0])
			b, eb := roman2int(strs[1])
			if ea != nil || eb != nil {
				panic(ERR_MIXED_NUMS)
			} else {
				res = int2roman(math_action(a, b, operator, true))
			}
		} else {
			res = strconv.Itoa(math_action(a, b, operator, false))
		}
		fmt.Println(RESULT_TXT + txt + " = " + res)
	}
}
