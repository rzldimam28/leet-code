package longestsubstringwithoutrepeatingcharacters

func LengthOfLongestSubstring(s string) int {
	m := make(map[string]int)

	for _, letter := range s {
		letterString := string(letter)
		_, ok := m[letterString]
		if !ok {
			m[letterString]++
		} else {
			break
		}
	}

	length := 0
	for range m {
		length++
	}

	return length
}