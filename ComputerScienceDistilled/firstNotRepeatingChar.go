package algorithms

// FirstNotRepeatingChar 找出字符串第一个只出现一次的字符
func FirstNotRepeatingChar(str string) byte {
	count := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		count[str[i]] = count[str[i]] + 1
	}
	for i := 0; i < len(str); i++ {
		if count[str[i]] == 1 {
			return str[i]
		}
	}

	return 0
}
