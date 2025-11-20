package stack

import "fmt"

type ValueT = string

type Stack struct { // naming
	storage  []ValueT // naming
	capacity int
}

func New(capacity int) *Stack {
	return &Stack{
		storage:  make([]ValueT, 0, capacity),
		capacity: capacity,
	}
}

// удаляет и возвращает
func (s *Stack) Pop() (ValueT, bool) {
	var val ValueT

	if len(s.storage) == 0 {
		return val, false
	}

	val = s.storage[len(s.storage)-1]

	s.storage = s.storage[:len(s.storage)-1]

	return val, true
}

func (s *Stack) Push(val ValueT) error {
	if len(s.storage) >= s.capacity {
		return fmt.Errorf("error")
	}
	s.storage = append(s.storage, val)

	return nil
}

func (s *Stack) Size() int {
	return len(s.storage)
}

func (s *Stack) String() (str string) {
	for i, el := range s.storage {
		str += fmt.Sprintf("индекс = %v, el = %v\n", i, el)
	}
	return
}
