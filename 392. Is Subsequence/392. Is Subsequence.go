func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	for i := range t {
		if len(s) == 0 {
			return true
		}
		if s[0] == t[i] {
			s = s[1:]
		}
	}
	return len(s) == 0
}