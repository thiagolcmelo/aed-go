package ds

import (
	"fmt"
	"strings"
)

type Set[K comparable] struct {
	data map[K]struct{}
}

func NewSet[K comparable]() *Set[K] {
	return &Set[K]{
		data: map[K]struct{}{},
	}
}

func (s *Set[K]) Add(v K) {
	s.data[v] = struct{}{}
}

func (s *Set[K]) Remove(v K) {
	delete(s.data, v)
}

func (s *Set[K]) Has(v K) bool {
	_, ok := s.data[v]
	return ok
}

func (s *Set[K]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Set[K]) Size() int {
	return len(s.data)
}

func (s *Set[K]) Clear() {
	s.data = map[K]struct{}{}
}

func (s *Set[K]) Elems() []K {
	elems := make([]K, 0)
	for e := range s.data {
		elems = append(elems, e)
	}
	return elems
}

func (s *Set[K]) Equal(o Set[K]) bool {
	for _, e := range s.Elems() {
		if !o.Has(e) {
			return false
		}
	}
	for _, e := range o.Elems() {
		if !s.Has(e) {
			return false
		}
	}
	return true
}

func (s *Set[K]) Intersection(o Set[K]) *Set[K] {
	i := NewSet[K]()
	for _, e := range s.Elems() {
		if o.Has(e) {
			i.Add(e)
		}
	}
	return i
}

func (s *Set[K]) Union(o Set[K]) *Set[K] {
	i := NewSet[K]()
	for _, e := range s.Elems() {
		i.Add(e)
	}
	for _, e := range o.Elems() {
		i.Add(e)
	}
	return i
}

func (s *Set[K]) Difference(o Set[K]) *Set[K] {
	i := NewSet[K]()
	for _, e := range s.Elems() {
		if !o.Has(e) {
			i.Add(e)
		}
	}
	return i
}

func (s *Set[K]) String() string {
	var elements []string
	for key := range s.data {
		elements = append(elements, fmt.Sprintf("%v", key))
	}
	return fmt.Sprintf("Set{%s}", strings.Join(elements, ", "))
}
