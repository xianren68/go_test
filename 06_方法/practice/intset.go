package practice

import (
	"bytes"
	"fmt"
)

// 定义bit数组
type IntSet struct {
	words []uint64
}

// 判断某个数是否存在（对应bit位为1）
func (s *IntSet) Has(x int) bool {
	// 获取这个值在集合中的索引及它在这个值的第几个bit位
	index, bit := x/64, uint(x%64)
	// 索引必须小于bit数组长度，并且值所在bit位为1
	return index < len(s.words) && s.words[index]&(1<<bit) != 0
}

// 往bit数组中添加一个值
func (s *IntSet) Add(x int) {
	index, bit := x/64, uint(x%64)
	// 如果要添加的值大于当前数组，则扩展数组
	for index >= len(s.words) {
		s.words = append(s.words, 0)
	}
	// 将对应位置修改为1
	s.words[index] |= 1 << bit
}

// 合并两个bit数组，相同位置并操作即可
func (s *IntSet) UnionWith(t *IntSet) {
	for index, value := range t.words {
		// 已存在此位置，对两个值进行|操作
		if index < len(s.words) {
			s.words[index] |= value
			// 不存在，则扩展数组
		} else {
			s.words = append(s.words, value)
		}
	}
}

// 打印有值的位置
func (s *IntSet) String() string {
	// byte buffer
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, val := range s.words {
		if val == 0 {
			continue
		}
		// 判断每一位是否存在
		for j := 0; j < 64; j++ {
			if val&(1<<uint(j)) != 0 {
				if buf.Len() > 1 {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
