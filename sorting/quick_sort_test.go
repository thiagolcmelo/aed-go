package sorting_test

import (
	"aed/sorting"
	"math/rand"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	testCases := []struct {
		input  []int
		output []int
	}{
		{input: []int{8, 1, 2, 3, 4, 7, 6, 5, 9, 0}, output: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{input: []int{}, output: []int{}},
		{input: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, output: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{input: []int{9, -8, 6, 7}, output: []int{-8, 6, 7, 9}},
	}

	for _, test := range testCases {
		sorting.QuickSort(test.input, 0, len(test.input)-1)
		if !reflect.DeepEqual(test.input, test.output) {
			t.Errorf("experado (%v), recebido (%v)", test.output, test.input)
		}
	}
}

func benchmarkQuickSort(n int, b *testing.B) {
	tests := make([][]int, b.N)
	for i := 0; i < b.N; i++ {
		tests[i] = make([]int, n)
		for j := 0; j < n; j++ {
			tests[i][j] = rand.Intn(100)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sorting.QuickSort(tests[i], 0, n-1)
	}
}

func benchmarkQuickSortInverse(n int, b *testing.B) {
	tests := make([][]int, b.N)
	for i := 0; i < b.N; i++ {
		tests[i] = make([]int, n)
		for j := 0; j < n; j++ {
			tests[i][j] = n - j
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sorting.QuickSort(tests[i], 0, n-1)
	}
}

func BenchmarkQuickSort1(b *testing.B) {
	benchmarkQuickSort(1, b)
}

func BenchmarkQuickSort2(b *testing.B) {
	benchmarkQuickSort(2, b)
}

func BenchmarkQuickSort3(b *testing.B) {
	benchmarkQuickSort(3, b)
}

func BenchmarkQuickSort4(b *testing.B) {
	benchmarkQuickSort(4, b)
}

func BenchmarkQuickSort5(b *testing.B) {
	benchmarkQuickSort(5, b)
}

func BenchmarkQuickSort6(b *testing.B) {
	benchmarkQuickSort(6, b)
}

func BenchmarkQuickSort7(b *testing.B) {
	benchmarkQuickSort(7, b)
}

func BenchmarkQuickSort8(b *testing.B) {
	benchmarkQuickSort(8, b)
}

func BenchmarkQuickSort9(b *testing.B) {
	benchmarkQuickSort(9, b)
}

func BenchmarkQuickSort10(b *testing.B) {
	benchmarkQuickSort(10, b)
}

func BenchmarkQuickSort20(b *testing.B) {
	benchmarkQuickSort(20, b)
}

func BenchmarkQuickSort30(b *testing.B) {
	benchmarkQuickSort(30, b)
}

func BenchmarkQuickSort40(b *testing.B) {
	benchmarkQuickSort(40, b)
}

func BenchmarkQuickSort50(b *testing.B) {
	benchmarkQuickSort(50, b)
}

func BenchmarkQuickSort60(b *testing.B) {
	benchmarkQuickSort(60, b)
}

func BenchmarkQuickSort70(b *testing.B) {
	benchmarkQuickSort(70, b)
}

func BenchmarkQuickSort80(b *testing.B) {
	benchmarkQuickSort(80, b)
}

func BenchmarkQuickSort90(b *testing.B) {
	benchmarkQuickSort(90, b)
}

func BenchmarkQuickSort100(b *testing.B) {
	benchmarkQuickSort(100, b)
}

func BenchmarkQuickSort200(b *testing.B) {
	benchmarkQuickSort(200, b)
}

func BenchmarkQuickSort300(b *testing.B) {
	benchmarkQuickSort(300, b)
}

func BenchmarkQuickSort400(b *testing.B) {
	benchmarkQuickSort(400, b)
}

func BenchmarkQuickSort500(b *testing.B) {
	benchmarkQuickSort(500, b)
}

func BenchmarkQuickSort600(b *testing.B) {
	benchmarkQuickSort(600, b)
}

func BenchmarkQuickSort700(b *testing.B) {
	benchmarkQuickSort(700, b)
}

func BenchmarkQuickSort800(b *testing.B) {
	benchmarkQuickSort(800, b)
}

func BenchmarkQuickSort900(b *testing.B) {
	benchmarkQuickSort(900, b)
}

func BenchmarkQuickSort1000(b *testing.B) {
	benchmarkQuickSort(1000, b)
}

func BenchmarkQuickSortInverse1(b *testing.B) {
	benchmarkQuickSortInverse(1, b)
}

func BenchmarkQuickSortInverse2(b *testing.B) {
	benchmarkQuickSortInverse(2, b)
}

func BenchmarkQuickSortInverse3(b *testing.B) {
	benchmarkQuickSortInverse(3, b)
}

func BenchmarkQuickSortInverse4(b *testing.B) {
	benchmarkQuickSortInverse(4, b)
}

func BenchmarkQuickSortInverse5(b *testing.B) {
	benchmarkQuickSortInverse(5, b)
}

func BenchmarkQuickSortInverse6(b *testing.B) {
	benchmarkQuickSortInverse(6, b)
}

func BenchmarkQuickSortInverse7(b *testing.B) {
	benchmarkQuickSortInverse(7, b)
}

func BenchmarkQuickSortInverse8(b *testing.B) {
	benchmarkQuickSortInverse(8, b)
}

func BenchmarkQuickSortInverse9(b *testing.B) {
	benchmarkQuickSortInverse(9, b)
}

func BenchmarkQuickSortInverse10(b *testing.B) {
	benchmarkQuickSortInverse(10, b)
}

func BenchmarkQuickSortInverse20(b *testing.B) {
	benchmarkQuickSortInverse(20, b)
}

func BenchmarkQuickSortInverse30(b *testing.B) {
	benchmarkQuickSortInverse(30, b)
}

func BenchmarkQuickSortInverse40(b *testing.B) {
	benchmarkQuickSortInverse(40, b)
}

func BenchmarkQuickSortInverse50(b *testing.B) {
	benchmarkQuickSortInverse(50, b)
}

func BenchmarkQuickSortInverse60(b *testing.B) {
	benchmarkQuickSortInverse(60, b)
}

func BenchmarkQuickSortInverse70(b *testing.B) {
	benchmarkQuickSortInverse(70, b)
}

func BenchmarkQuickSortInverse80(b *testing.B) {
	benchmarkQuickSortInverse(80, b)
}

func BenchmarkQuickSortInverse90(b *testing.B) {
	benchmarkQuickSortInverse(90, b)
}

func BenchmarkQuickSortInverse100(b *testing.B) {
	benchmarkQuickSortInverse(100, b)
}

func BenchmarkQuickSortInverse200(b *testing.B) {
	benchmarkQuickSortInverse(200, b)
}

func BenchmarkQuickSortInverse300(b *testing.B) {
	benchmarkQuickSortInverse(300, b)
}

func BenchmarkQuickSortInverse400(b *testing.B) {
	benchmarkQuickSortInverse(400, b)
}

func BenchmarkQuickSortInverse500(b *testing.B) {
	benchmarkQuickSortInverse(500, b)
}

func BenchmarkQuickSortInverse600(b *testing.B) {
	benchmarkQuickSortInverse(600, b)
}

func BenchmarkQuickSortInverse700(b *testing.B) {
	benchmarkQuickSortInverse(700, b)
}

func BenchmarkQuickSortInverse800(b *testing.B) {
	benchmarkQuickSortInverse(800, b)
}

func BenchmarkQuickSortInverse900(b *testing.B) {
	benchmarkQuickSortInverse(900, b)
}

func BenchmarkQuickSortInverse1000(b *testing.B) {
	benchmarkQuickSortInverse(1000, b)
}
