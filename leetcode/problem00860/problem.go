package problem00860

func lemonadeChange(bills []int) bool {
	fiveCount, tenCount := 0, 0
	for _, bill := range bills {
		switch bill {
		case 5:
			fiveCount++
		case 10:
			if fiveCount == 0 {
				return false
			}
			fiveCount--
			tenCount++
		default: // 20:
			// 优先找 10 块的
			if tenCount > 0 {
				if fiveCount == 0 {
					return false
				}
				fiveCount--
				tenCount--
			} else {
				if fiveCount < 3 {
					return false
				}
				fiveCount -= 3
			}
		}
	}
	return true
}
