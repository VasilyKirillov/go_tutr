package datastruct

import (
	"fmt"
)

type stack []interface{}

func NewStack(elems ...interface{}) stack {
	if elems == nil {
		var s []interface{}
		return s
	}
	s := make([]interface{}, len(elems))
	copy(s, elems)
	return s
}

func (s *stack) Pop() interface{} {
	size := len(*s) - 1
	elem := (*s)[size]
	(*s)[size] = nil
	*s = (*s)[:size]
	return elem
}

func (s *stack) Push(elem interface{}) {
	*s = append(*s, elem)
}

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *stack) Size() int {
	return len(*s)
}

func (s stack) String() string {
	size := s.Size()
	var res string = "["
	for i, el := range s {
		res += fmt.Sprint(el)
		if i < size-1 {
			res += ","
		}
	}
	return res + "]"
}
