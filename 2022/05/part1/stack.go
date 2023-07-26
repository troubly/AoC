package main

import "errors"

type stack[T any] struct {
	s []T
}

func (s *stack[T]) push(v T) {
	s.s = append(s.s, v)
}

func (s *stack[T]) pop() (T, error) {
	if len(s.s) == 0 {
		var empty T
		return empty, errors.New("stack is empty")
	}

	ret := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]

	return ret, nil
}
