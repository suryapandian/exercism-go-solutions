package luhn

// Valid checks if the input string is valid luhn
func Valid(input string) bool {
	sum, count := luhnSum([]rune(input))

	return count > 1 && sum%10 == 0
}

func luhnSum(runes []rune) (sum, count int32) {
	for i := len(runes) - 1; i >= 0; i-- {
		r := runes[i]
		if r == ' ' {
			continue
		}
		if r < '0' || r > '9' {
			return 0, 0
		}

		sum += modVal(r-'0', count)
		count++
	}
	return sum, count
}

func modVal(value, count int32) int32 {
	if count%2 != 0 {
		value *= 2
		if value > 9 {
			value -= 9
		}
	}
	return value
}
