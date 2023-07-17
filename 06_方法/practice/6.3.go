package practice

// 返回所有数据
func (s *IntSet) Elems() (elems []int) {
	for i, val := range s.words {
		if val == 0 {
			continue
		}
		for j := 0; j < ADAPTATION; j++ {
			if val&uint(1<<j) != 0 {
				elems = append(elems, i*ADAPTATION+j)
			}
		}
	}
	return
}
