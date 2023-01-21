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

var INPUT = "Input:\n"
var OUTPUT = "Output:\n"
var WARNING = "[WARNING]"
var ERROR = "[ERROR]"
var RESULT = "Result is"

var isLatin = false
var isArabic = false

func main() {
	var operand string

	var f *os.File
	f = os.Stdin
	defer f.Close()
	scanner := bufio.NewScanner(f)
	fmt.Println(INPUT)
	for scanner.Scan() {

		str := scanner.Text()
		args := strings.Split(str, " ")

		val1, err := parseIntArgument(args[0])
		operand, err = parseOperandArgument(args[1])
		val2, err := parseIntArgument(args[2])
		if err {
			fmt.Println(OUTPUT, ERROR, "Programm exit with error")
			os.Exit(1)
		}

		if isLatin && isArabic {
			fmt.Println(OUTPUT, ERROR, "Latins and arabic numbers not supported together")
			os.Exit(1)
		}

		var val int
		switch operand {
		case "Plus":
			val = plus(val1, val2)
			break
		case "Minus":
			val = minus(val1, val2)
			break
		case "Multiple":
			val = multiple(val1, val2)
			break
		case "Divide":
			val = divide(val1, val2)
			break
		}

		if isLatin {
			fmt.Println(OUTPUT, RESULT, arabToLatin(val))
		} else {
			fmt.Println(OUTPUT, RESULT, val)
		}
		os.Exit(0)
	}
}

func plus(val1 int, val2 int) int {
	return val1 + val2
}

func minus(val1 int, val2 int) int {
	return val1 - val2
}

func multiple(val1 int, val2 int) int {
	return val1 * val2
}

func divide(val1 int, val2 int) int {
	return val1 / val2
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
			fmt.Println(WARNING, "Argument is not a integer", error)
			return 0, true
		}
		return firstInt, false
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
