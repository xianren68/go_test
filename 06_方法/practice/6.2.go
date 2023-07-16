package practice

// 添加多个值
func (s *IntSet) AddAll(set ...int) {
	for _, val := range set {
		s.Add(val)
	}
}

// 交集
func (s *IntSet) IntsectWith(t *IntSet) {
	for i, val := range s.words {
		if val == 0 {
			continue
		}
		// 这个值不越界且这个位置不为0
		if i < len(t.words) && t.words[i] != 0 {
			s.words[i] &= t.words[i]
			// 越界/为0
		} else {
			s.words[i] = 0
		}
	}
}

// 差集
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, val := range s.words {
		// 当前值为0
		if val == 0 {
			continue
		}
		// 长度超出时/另一个集合对应位置为0，不用变
		if i >= len(t.words) || t.words[i] == 0 {
			continue
		}
		s.words[i] = val ^ t.words[i]

	}
}

// 并差集(两个集合的并集-交集)
func (s *IntSet) SymmetricDifference(t *IntSet) {
	// 一个新的集合
	u := s.Copy()
	// 并集
	u.UnionWith(t)
	// 交集
	s.IntsectWith(t)
	u.DifferenceWith(s)
	s.words = u.words

}
