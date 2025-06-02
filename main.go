package main

import "fmt"

func num2geez(n int) string {
	numberMap := map[int]rune{
		1: '፩', 2: '፪', 3: '፫', 4: '፬', 5: '፭',
		6: '፮', 7: '፯', 8: '፰', 9: '፱', 10: '፲',
		20: '፳', 30: '፴', 40: '፵', 50: '፶', 60: '፷',
		70: '፸', 80: '፹', 90: '፺', 100: '፻', 10000: '፼',
	}

	res := ""
	subscript := 0

	for n > 0 {
		value := n % 100

		ones := value % 10
		tens := value / 10 * 10

		if subscript > 0 && !(tens == 0 && ones == 0) {
			if subscript%2 == 0 {
				res = string(numberMap[10000]) + res
			} else {
				res = string(numberMap[100]) + res
			}
			if tens == 0 && ones == 1 {
				ones = 0
			}
		}

		res = string(numberMap[tens]) + string(numberMap[ones]) + res

		n /= 100
		subscript++

	}

	return res
}

func main() {
	// for i := 1; i <= 100; i++ {
	// 	fmt.Println(i, num2geez(i))
	// }
	for i := 1; i <= 1000000; i *= 10 {
		fmt.Println(i, num2geez(i))
	}

	// fmt.Println(num2geez(170))
	// fmt.Println(num2geez(200))
	// fmt.Println(num2geez(1170))
	// fmt.Println(num2geez(9999))
}
