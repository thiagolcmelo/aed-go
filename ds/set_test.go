package ds_test

import (
	"aed/ds"
	"testing"
)

func TestSetAddHas(t *testing.T) {
	testCases := []struct {
		elems []int
	}{
		{elems: []int{}},
		{elems: []int{10, 19, 39, 12, 10}},
	}

	for _, test := range testCases {
		s := ds.NewSet[int]()
		for _, e := range test.elems {
			s.Add(e)
		}

		for _, e := range test.elems {
			if !s.Has(e) {
				t.Errorf("Should have %d, but it is missing", e)
			}
		}
	}
}

func TestSetAddRemove(t *testing.T) {
	testCases := []struct {
		add    []int
		remove []int
		left   []int
	}{
		{add: []int{10, 19, 39, 12, 10}, remove: []int{10}, left: []int{19, 39, 12}},
	}

	for _, test := range testCases {
		s := ds.NewSet[int]()
		for _, e := range test.add {
			s.Add(e)
		}
		for _, e := range test.remove {
			s.Remove(e)
		}
		for _, e := range test.left {
			if !s.Has(e) {
				t.Errorf("Should have %d, but it is missing", e)
			}
		}
		for _, e := range test.remove {
			if s.Has(e) {
				t.Errorf("Should have removed %d, but it is still there", e)
			}
		}
	}
}

func TestSetIntersection(t *testing.T) {
	testCases := []struct {
		sideA        []int
		sideB        []int
		intersection []int
	}{
		{
			sideA:        []int{1, 2, 3, 4, 5},
			sideB:        []int{3, 4, 5, 6, 7},
			intersection: []int{3, 4, 5},
		},
		{
			sideA:        []int{1},
			sideB:        []int{3},
			intersection: []int{},
		},
		{
			sideA:        []int{1, 2, 3, 4, 5},
			sideB:        []int{},
			intersection: []int{},
		},
		{
			sideA:        []int{},
			sideB:        []int{1, 2, 3, 4, 5},
			intersection: []int{},
		},
	}

	for _, test := range testCases {
		a := ds.NewSet[int]()
		for _, e := range test.sideA {
			a.Add(e)
		}

		b := ds.NewSet[int]()
		for _, e := range test.sideB {
			b.Add(e)
		}

		i := ds.NewSet[int]()
		for _, e := range test.intersection {
			i.Add(e)
		}

		r := a.Intersection(*b)
		if !r.Equal(*i) {
			t.Errorf("Expected %s, got %s", i.String(), r.String())
		}
	}
}
