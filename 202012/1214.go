package leetcode

//https://leetcode-cn.com/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string, len(strs))
	result := make([][]string, 0, len(strs))
	for _, v := range strs {
		alphabet := [26]int{}
		for _, w := range []byte(v) {
			alphabet[w-'a']++
		}
		m[alphabet] = append(m[alphabet], v)
	}
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

/* time limits
func alphabet2string(alphabet []int) string {
	b := strings.Builder{}
	for _, v := range alphabet {
		b.WriteString(strconv.Itoa(v))
	}
	return b.String()
}
func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0, len(strs))
	m := make([][]int, len(strs))
	for i, v := range strs {
		alphabet := make([]int, 26)
		for _, w := range []byte(v) {
			alphabet[w-'a']++
		}
		m[i] = alphabet
	}
	for i := 0; i < len(strs); i++ {
		if len(m[i]) == 0 {
			continue
		}
		tmp := []string{strs[i]}
		for j := i + 1; j < len(strs); j++ {
			if len(m[j]) == 0 {
				continue
			}
			if reflect.DeepEqual(m[i], m[j]) {
				tmp = append(tmp, strs[j])
				m[j] = []int{}
			}
		}
		result = append(result, tmp)
	}
	return result
}
*/
