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
