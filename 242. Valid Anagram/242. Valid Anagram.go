func isAnagram(s string, t string) bool {
	sCount := make(map[byte]int, 26)
	tCount := make(map[byte]int, 26)
	if len(s) != len(t) {
		return false
	}
	for i := range s {
		sCount[s[i]] += 1
		tCount[t[i]] += 1
	}
	for i := 0; i < 26; i++ {
		if sCount[byte(i+97)] != tCount[byte(i+97)] {
			return false
		}
	}
	return true
}