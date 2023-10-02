package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	for {
		fmt.Println("Калькулятор римских и арабских чисел.")
		fmt.Println("Введите выражение (например, II+III или 2+3), или введите 'stop' для завершения)")

		var input string
		fmt.Scanln(&input)
		strings.ReplaceAll(input, " ", "")
		if input == "stop" {
			fmt.Println("Программа завершена.")
			return
		}
		operatorPos := strings.IndexAny(input, "+-*/")
		if operatorPos < 1 || operatorPos == len(input)-1 {
			fmt.Println("Неверный формат ввода.")
			continue
		}
		num1 := input[:operatorPos]
		operator := string(input[operatorPos])
		num2 := input[operatorPos+1:]

		isArabic1 := isArabic(num1)
		isArabic2 := isArabic(num2)
		isRoman1 := isRoman(num1)
		isRoman2 := isRoman(num2)

		if (isArabic1 && isRoman2) || (isRoman1 && isArabic2) {
			fmt.Println("Ошибка: Введены неправильные данные (либо арабские, либо римские числа I-X)")
			continue
		}

		var result string
		if isArabic1 && isArabic2 {
			a, b := parseArabic(num1), parseArabic(num2)
			result = performArabicOperation(a, operator, b)
		} else if isRoman1 && isRoman2 {
			a, b := romanToArabic(num1), romanToArabic(num2)
			arabicResult := performArabicOperation(a, operator, b)
			if arabicResult != "Ошибка: Неподдерживаемая операция" {
				result = arabicToRoman(parseArabic(arabicResult))
			} else {
				result = arabicResult
			}

		} else {
			fmt.Println("Ошибка: Введены неправильные данные (либо арабские, либо римские числа I-X)")
			continue
		}
		fmt.Printf("Результат: %s\n", result)

	}

}

func isArabic(input string) bool {
	for _, char := range input {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func isRoman(input string) bool {
	validRomanNumerals := map[string]bool{
		"I": true, "II": true, "III": true, "IV": true, "V": true,
		"VI": true, "VII": true, "VIII": true, "IX": true, "X": true,
	}
	return validRomanNumerals[input]
}

func parseArabic(input string) int {
	var result int
	fmt.Sscanf(input, "%d", &result)
	return result
}

func performArabicOperation(num1 int, operator string, num2 int) string {
	switch operator {
	case "+":
		return fmt.Sprintf("%d", num1+num2)
	case "-":
		return fmt.Sprintf("%d", num1-num2)
	case "*":
		return fmt.Sprintf("%d", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Ошибка: Деление на ноль.")
		}
		return fmt.Sprintf("%d", num1/num2)
	default:
		return "Ошибка: Неподдерживаемая операция"

	}

}

func romanToArabic(input string) int {
	romanNumerals := map[rune]int{
		'I': 1, 'V': 5, 'X': 10,
	}
	var result int
	prevValue := 0
	for i := len(input) - 1; i >= 0; i-- {
		value := romanNumerals[rune(input[i])]
		if value < prevValue {
			result -= value
		} else {
			result += value
		}
		prevValue = value
	}
	return result
}

func arabicToRoman(input int) string {
	if input <= 0 {
		return "Ошибка: Результат не может быть меньше или равен нулю"
	}

	romanNumerals := []struct {
		Value  int
		Symbol string
	}{
		{10, "X"}, {9, "IX"}, {5, "V"},
		{4, "IV"}, {1, "I"},
	}
	var result strings.Builder
	for _, numeral := range romanNumerals {
		for input >= numeral.Value {
			result.WriteString(numeral.Symbol)
			input -= numeral.Value

		}
	}
	return result.String()
}
