package geeznum

// NumToGeez converts an integer to a Geez numeral string.
func NumToGeez(n int) string {
	numberMap := map[int]rune{
		1: '፩', 2: '፪', 3: '፫', 4: '፬', 5: '፭',
		6: '፮', 7: '፯', 8: '፰', 9: '፱', 10: '፲',
		20: '፳', 30: '፴', 40: '፵', 50: '፶', 60: '፷',
		70: '፸', 80: '፹', 90: '፺', 100: '፻', 10000: '፼',
	}

	result := ""
	place := 0

	for n > 0 {
		value := n % 100
		ones := value % 10
		tens := value / 10 * 10
		sep := 0

		if place != 0 {
			if place%2 == 0 {
				sep = 10000
			} else if ones != 0 || tens != 0 {
				sep = 100
			}
			result = string(numberMap[sep]) + result
		}

		if ones == 1 && tens == 0 && place > 0 {
			if sep == 100 || n/100 == 0 {
				ones = 0
			}
		}

		if ones != 0 {
			result = string(numberMap[ones]) + result
		}

		if tens != 0 {
			result = string(numberMap[tens]) + result
		}

		n /= 100
		place++
	}

	return result
}
