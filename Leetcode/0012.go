package leetcoode

var valueSymbols = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

var values = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

func intToRoman(num int) string {
	var result string
	for num > 0 {
		i := 12
		for ; i >= 0; i-- {
			if num >= values[i] {
				break
			}
		}
		// no negative detect
		result += valueSymbols[values[i]]
		num -= values[i]
	}
	return result
}
