package computers_count

import "strconv"

var computers = []string{"компьютер", "компьютера", "компьютеров"}

// ComputersCount возвращает строку вида "N компьютер", "N компьютера", "N компьютеров"
func ComputersCount(count int) string {
	var strCount = strconv.Itoa(count)
	var result = strCount + " "

	// 5-20 компьютеров
	if count >= 5 && count <= 20 || count <= -5 && count >= -20 {
		return result + computers[2]
	}

	// Последняя цифра
	switch strCount[len(strCount)-1:] {
	case "1":
		result += computers[0]
	case "2", "3", "4":
		result += computers[1]
	case "5", "6", "7", "8", "9", "0":
		result += computers[2]
	}

	return result
}
