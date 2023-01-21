package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Operands struct {
	Plus     string
	Minus    string
	Multiply string
	Divide   string
}

var operands = Operands{
	"+",
	"-",
	"*",
	"/",
}

var romans = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

var MIN_ARGS_LENGTH = 3
var MAX = 10
var MIN = 1
var INPUT = "Input:"
var OUTPUT = "Output:\n"
var WARNING = "[WARNING]"
var ERROR = "[ERROR]"
var RESULT = "Result is"
var EXIT_WITH_ERROR = "Programm exit with error"

var isLatin = false
var isArabic = false

func main() {

	var f *os.File
	f = os.Stdin
	defer f.Close()
	scanner := bufio.NewScanner(f)
	fmt.Println(INPUT)
	for scanner.Scan() {
		str := scanner.Text()
		args := strings.Split(str, " ")

		length := len(args)
		if length < MIN_ARGS_LENGTH {
			fmt.Printf("%s %s Expected number of arguments '%d', received '%d'\n", OUTPUT, ERROR, MIN_ARGS_LENGTH, length)
			os.Exit(1)
		}

		arg1, arg2, errN := parseNumbers(args)
		operand, errO := parseOperandArgument(args[1])
		if errN || errO {
			fmt.Println(OUTPUT, ERROR, EXIT_WITH_ERROR)
			os.Exit(1)
		}

		val := operate(arg1, arg2, operand)
		showResult(val)
		creanGlobal()

		fmt.Println(INPUT)
	}
}

func creanGlobal() {
	isLatin = false
	isArabic = false
}

func showResult(val int) {
	if isLatin {
		if val < MIN {
			fmt.Printf("%s The result cannot be converted to latin: is '%d' that is lower than '%d'\n", OUTPUT, val, MIN)
			return
		}
		fmt.Println(OUTPUT, RESULT, arabToLatin(val))
	} else {
		fmt.Println(OUTPUT, RESULT, val)
	}
}

func parseNumbers(args []string) (int, int, bool) {
	arg1, err1 := parseIntArgument(args[0])
	arg2, err2 := parseIntArgument(args[2])
	if err1 || err2 {
		return 0, 0, true
	}

	if isLatin && isArabic {
		fmt.Println(WARNING, "Latins and arabic numbers not supported together")
		return 0, 0, true
	}
	return arg1, arg2, false
}

func plus(arg1 int, arg2 int) int {
	return arg1 + arg2
}

func minus(arg1 int, arg2 int) int {
	return arg1 - arg2
}

func multiple(arg1 int, arg2 int) int {
	return arg1 * arg2
}

func divide(arg1 int, arg2 int) int {
	return arg1 % arg2
}

func operate(arg1 int, arg2 int, operand string) int {
	var val int
	switch operand {
	case "Plus":
		val = plus(arg1, arg2)
		break
	case "Minus":
		val = minus(arg1, arg2)
		break
	case "Multiple":
		val = multiple(arg1, arg2)
		break
	case "Divide":
		val = divide(arg1, arg2)
		break
	}
	return val
}

func checkRange(num int) bool {
	if num < MIN || num > MAX {
		fmt.Printf("%s Invalid number '%d' expected: [1 - 10 or I - X]\n", WARNING, num)
		return false
	}
	return true
}

func parseIntArgument(arg string) (int, bool) {
	firstInt, err := strconv.Atoi(arg)
	if err == nil {
		isArabic = true
	} else {
		firstInt, ok := romans[arg]
		if ok {
			isLatin = true
		} else {
			error := fmt.Errorf("%w", err)
			fmt.Println(WARNING, "Argument is not a integer number", error)
			return 0, true
		}
		if !checkRange(firstInt) {
			return 0, true
		}
		return firstInt, false
	}
	if !checkRange(firstInt) {
		return 0, true
	}
	return firstInt, false
}

func parseOperandArgument(arg string) (string, bool) {
	v := reflect.ValueOf(operands)
	var name string
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i)
		if value.String() == arg {
			name = v.Type().Field(i).Name
		}
	}
	if name == "" {
		fmt.Println(WARNING, "Operand is missing", name)
		return name, true
	}
	return name, false
}

func arabToLatin(arg int) string {
	var value string
	for key, val := range romans {
		if val == arg {
			value = key
		}
	}
	return value
}
