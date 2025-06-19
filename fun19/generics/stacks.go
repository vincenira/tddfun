package generics

type StacksOfInts struct {
	values []int
}

func (s *StacksOfInts) Push(value int) {
	s.values = append(s.values, value)
}

func (s *StacksOfInts) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *StacksOfInts) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	index := len(s.values) - 1
	lastValue := s.values[index]
	s.values = s.values[:index]
	return lastValue, true
}

type StacksOfStrings struct {
	values []string
}

func (s *StacksOfStrings) Push(value string) {
	s.values = append(s.values, value)
}

func (s *StacksOfStrings) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *StacksOfStrings) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	index := len(s.values) - 1
	lastValue := s.values[index]
	s.values = s.values[:index]
	return lastValue, true
}

type (
	StackOfInts    = Stack
	StackOfStrings = Stack
)

type Stack struct {
	values []interface{}
}

func (s *Stack) Push(value interface{}) {
	s.values = append(s.values, value)
}

func (s *Stack) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		var zero interface{}
		return zero, false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]
	return el, true
}

type StackG[T any] struct {
	values []T
}

// Making a constructor for Stack[T]
func NewStack[T any]() *StackG[T] {
	return new(StackG[T])
}

func (s *StackG[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *StackG[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *StackG[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]
	return el, true
}
