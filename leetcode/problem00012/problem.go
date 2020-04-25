package problem00012

import "strings"

func intToRomanStaticTable(num int) string {
	romanTable := [][]string{
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}, // 个位
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}, // 十位
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}, // 百位
		{"", "M", "MM", "MMM"}, // 千位
	}

	return romanTable[3][num/1000] + romanTable[2][num/100%10] + romanTable[1][num/10%10] + romanTable[0][num%10]
}

func intToRomanGreedyAlgorithms(num int) string {
	romanTable := []struct {
		num   int
		roman string
	}{
		{num: 1000, roman: "M"},
		{num: 900, roman: "CM"},
		{num: 500, roman: "D"},
		{num: 400, roman: "CD"},
		{num: 100, roman: "C"},
		{num: 90, roman: "XC"},
		{num: 50, roman: "L"},
		{num: 40, roman: "XL"},
		{num: 10, roman: "X"},
		{num: 9, roman: "IX"},
		{num: 5, roman: "V"},
		{num: 4, roman: "IV"},
		{num: 1, roman: "I"},
	}

	ret := strings.Builder{}
	for idx := 0; idx < len(romanTable); idx++ {
		for num >= romanTable[idx].num {
			ret.WriteString(romanTable[idx].roman)
			num -= romanTable[idx].num
		}
	}

	return ret.String()
}
