package main

import (
	"fmt"
	"strconv"
)

func convert_num_to_words_100s(num int) string {

	var one_digit = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var one_and_a_digit = []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	var two_digits = []string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

	str_num := strconv.Itoa(num)
	length := len(str_num)
	if length == 0 {
		return ("Zero")
	} else if length == 1 {
		var int_num = int(str_num[0] - '0')
		return (one_digit[int_num])
	} else if length == 2 {
		if str_num[0] == '1' {
			var int_num = int(str_num[1] - '0')
			return (one_and_a_digit[int_num])
		} else if str_num[1] == '0' {
			var int_first_digit = int(str_num[0]-'0') - 2
			return (two_digits[int_first_digit])
		} else {
			var int_first_digit = int(str_num[0]-'0') - 2
			var int_second_digit = int(str_num[1] - '0')
			return (two_digits[int_first_digit] + " " + one_digit[int_second_digit])
		}
	} else if length == 3 {
		var int_first_digit = int(str_num[0] - '0')
		if str_num[1] == '0' {
			var int_third_digit = int(str_num[2] - '0')
			if str_num[2] == '0' {
				return (one_digit[int_first_digit] + " Hundret ")
			} else {
				return (one_digit[int_first_digit] + " Hundret and " + one_digit[int_third_digit])
			}
		} else if str_num[1] == '1' {
			var int_second_and_third_digit = int(str_num[2] - '0')
			return (one_digit[int_first_digit] + " Hundret and " + one_and_a_digit[int_second_and_third_digit])
		} else {
			var int_second_digit = int(str_num[1]-'0') - 2
			var int_third_digit = int(str_num[2] - '0')
			if int_third_digit == 0 {
				return (one_digit[int_first_digit] + " Hundret and " + two_digits[int_second_digit])
			} else {
				return (one_digit[int_first_digit] + " Hundret and " + two_digits[int_second_digit] + " " + one_digit[int_third_digit])
			}
		}
	}

	return ("")
}

func convert_num_to_words(num int) {

	str_num := strconv.Itoa(num)
	length := len(str_num)

	if length == 4 {

		var one_digit = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

		intVar, _ := strconv.Atoi(str_num[1:])
		str := convert_num_to_words_100s(intVar)

		first_digit := int(str_num[0] - '0')

		if str_num[1] == '0' && str_num[2] == '0' && str_num[3] == '0' {
			fmt.Println(one_digit[first_digit] + " Thousand")
		} else {
			fmt.Println(one_digit[first_digit] + " Thousand and " + str)
		}
	} else {
		str := convert_num_to_words_100s(num)
		fmt.Println(str)
	}
}

func main() {
	fmt.Println("Enter your Number?")
	var input int
	fmt.Scanln(&input)
	convert_num_to_words(input)
}
