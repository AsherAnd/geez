package main

import "fmt"

func num2geez(n int) string {
	numberMap := map[int]rune{1: '፩', 2: '፪', 3: '፫', 4: '፬', 5: '፭', 6: '፮', 7: '፯', 8: '፰', 9: '፱',
		10: '፲', 20: '፳', 30: '፴', 40: '፵', 50: '፶', 60: '፷', 70: '፸', 80: '፹', 90: '፺',
		100: '፻', 10000: '፼'}

	res := ""
	place := 1

	for n > 0 {
		value := n % 10

		if value != 0 {
			if ((place%100 == 0) || place%10000 == 0) && value > 1 {
				res = string(numberMap[place]) + res
				place = 1
			}

			res = string(numberMap[value*place]) + res

		}

		n /= 10
		place *= 10

	}

	return res
}

func main() {
	for i := 1; i <= 1000; i++ {
		fmt.Println(i, num2geez(i))
	}
	// for i := 1; i <= 1000000; i *= 10 {
	// 	fmt.Println(i, num2geez(i))
	// }
}
