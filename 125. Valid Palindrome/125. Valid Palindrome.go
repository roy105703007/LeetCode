func isPalindrome(s string) bool {
	data := []byte(s)
	left := 0
	right := len(s) - 1
	for left < right {
		if data[left] >= 65 && data[left] <= 90 {
			data[left] = data[left] + 32
		} else if (data[left] < 97 || data[left] > 122) && (data[left] < 48 || data[left] > 57) {
			left++
			continue
		}
		if data[right] >= 65 && data[right] <= 90 {
			data[right] = data[right] + 32
		} else if (data[right] < 97 || data[right] > 122) && (data[right] < 48 || data[right] > 57) {
			right--
			continue
		}
		if data[left] != data[right] {
			return false
		}
		left++
		right--
	}
	return true
}