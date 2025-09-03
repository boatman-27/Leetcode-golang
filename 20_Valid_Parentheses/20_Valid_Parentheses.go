package main

import "fmt"

type Node struct {
	parentheses rune
	Next        *Node
}

type Stack struct {
	head   *Node
	length int
}

func (s *Stack) Push(data rune) {
	newNode := &Node{parentheses: data, Next: s.head}
	s.head = newNode
	s.length++
}

func (s *Stack) Pop() {
	poppedNode := s.head
	s.head = poppedNode.Next
	s.length--
}

func (s *Stack) Peek() rune {
	return s.head.parentheses
}

func (s *Stack) Print() {
	current := s.head
	for current != nil {
		fmt.Printf("%c -> ", current.parentheses)
		current = current.Next
	}
	fmt.Println("nil")
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	paren := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := &Stack{}
	for _, r := range s {
		if _, ok := paren[r]; !ok {

			if stack.length == 0 {
				return false
			}
			top := stack.Peek()
			if string(paren[top]) != string(r) {
				return false
			}
			stack.Pop()
		} else {
			stack.Push(r)
		}
	}
	if stack.length == 0 {
		return true
	} else {
		return false
	}
}

// func main() {
// 	s := "(("
// 	fmt.Println(isValid(s))
// }
