package practice

// 返回bit数组中元素的个数
func (s *IntSet) Len() (count int) {
	for _, val := range s.words {
		if val == 0 {
			continue
		}
		for j := 0; j < ADAPTATION; j++ {
			if val&(1<<uint(j)) != 0 {
				count++
			}
		}

	}
	return

}

// 删除bit数组中的某个元素
func (s *IntSet) Remove(x int) {
	index, bit := x/ADAPTATION, uint(x%ADAPTATION)
	if index >= len(s.words) {
		return
	}
	// 将对应位置置零
	s.words[index] &= ^(1 << bit)
}

// 清空bit数组
func (s *IntSet) Clear() {
	s.words = nil
}

// 将当前bit数组copy
func (s *IntSet) Copy() *IntSet {
	return &IntSet{append([]uint{}, s.words...)}
}
