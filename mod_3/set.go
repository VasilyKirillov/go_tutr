package datastruct

import "fmt"

type set struct {
	holder map[interface{}]bool
}

func NewSet() set {
	s := make(map[interface{}]bool)
	return set{holder: s}
}

func (s *set) Add(el interface{}) bool {
	_, ok := s.holder[el]
	if ok {
		return false
	}
	s.holder[el] = true
	return true
}

func (s *set) Remove(el interface{}) bool {
	_, ok := s.holder[el]
	if ok {
		delete(s.holder, el)
		return true
	}
	return false
}

func (s *set) Contains(target interface{}) bool {
	for el, _ := range s.holder {
		if el == target {
			return true
		}
	}
	return false
}

func (s set) Iterator() chan interface{} {
	size := len(s.holder)
	ch := make(chan interface{}, size-1)
	go func() {
		for k := range s.holder {
			ch <- k
		}
		close(ch)
	}()
	return ch
}

func (s set) Size() int {
	return len(s.holder)
}

func (s set) String() string {
	i := 0
	size := s.Size()
	res := "["
	for k := range s.holder {
		i++
		res += fmt.Sprint(k)
		if i < size {
			res += ","
		}

	}
	return res + "]"
}
