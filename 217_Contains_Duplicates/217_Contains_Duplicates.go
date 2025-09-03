package main

type Set struct {
	set map[int]struct{}
}

func newSet() *Set {
	return &Set{
		set: make(map[int]struct{}),
	}
}

func (s *Set) AddElement(val int) {
	s.set[val] = struct{}{}
}

func (s *Set) CheckExists(val int) bool {
	_, ok := s.set[val]
	return ok
}

func containsDuplicate(nums []int) bool {
	set := newSet()
	for _, value := range nums {
		if set.CheckExists(value) {
			return true
		}

		set.AddElement(value)
	}

	return false
}

// func main() {
// 	nums := []int{1, 2, 3, 4}
// 	fmt.Println(containsDuplicate(nums))
// }
