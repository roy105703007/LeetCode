func partitionLabels(s string) []int {
	count := make(map[byte]int, 26)
	res := []int{}
	last := 0
	for i := range s {
		count[s[i]-'a'] = i
	}
	end := count[s[0]-'a']
	fmt.Println(count)
	for i := range s {
		if i == end {
			res = append(res, i-last+1)
			if i == len(s)-1 {
				return res
			}
			last = i + 1
			end = count[s[last]-'a']
		} else if count[s[i]-'a'] > end {
			end = count[s[i]-'a']
		}
	}
	return res
}