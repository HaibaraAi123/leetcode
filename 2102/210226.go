package leetcode

//https://leetcode-cn.com/problems/number-of-valid-words-for-each-puzzle/
func findNumOfValidWords(words []string, puzzles []string) []int {
	res := make([]int, len(puzzles))
	puzzlesMAP := make([]map[int]bool, len(puzzles))
	for i, v := range puzzles {
		puzzlesMAP[i] = make(map[int]bool, 7)
		for _, newv := range v {
			puzzlesMAP[i][int(newv-'a')] = true
		}
	}
	for _, v := range words {
		tmp := make(map[int]bool, 26)
		for _, newv := range v {
			tmp[int(newv-'a')] = true
		}
		for i := range puzzles {
			if _, ok := tmp[int(puzzles[i][0]-'a')]; !ok {
				continue
			} else {
				flag := false
				for j, ok := range tmp {
					if ok {
						if puzzlesMAP[i][j] == false {
							flag = true
							break
						}
					}
				}
				if flag == false {
					res[i]++
				}
				continue
			}
		}
	}
	return res
}
