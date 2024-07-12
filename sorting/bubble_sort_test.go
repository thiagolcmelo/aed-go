package sorting_test

import (
	"aed/sorting"
	"math/rand"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
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
		sorting.BubbleSort(test.input)
		if !reflect.DeepEqual(test.input, test.output) {
			t.Errorf("experado (%v), recebido (%v)", test.output, test.input)
		}
	}
}

func benchmarkBubbleSort(n int, b *testing.B) {
	tests := make([][]int, b.N)
	for i := 0; i < b.N; i++ {
		tests[i] = make([]int, n)
		for j := 0; j < n; j++ {
			tests[i][j] = rand.Intn(100)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sorting.BubbleSort(tests[i])
	}
}

func BenchmarkBubbleSort1(b *testing.B) {
	benchmarkBubbleSort(1, b)
}

func BenchmarkBubbleSort2(b *testing.B) {
	benchmarkBubbleSort(2, b)
}

func BenchmarkBubbleSort3(b *testing.B) {
	benchmarkBubbleSort(3, b)
}

func BenchmarkBubbleSort4(b *testing.B) {
	benchmarkBubbleSort(4, b)
}

func BenchmarkBubbleSort5(b *testing.B) {
	benchmarkBubbleSort(5, b)
}

func BenchmarkBubbleSort6(b *testing.B) {
	benchmarkBubbleSort(6, b)
}

func BenchmarkBubbleSort7(b *testing.B) {
	benchmarkBubbleSort(7, b)
}

func BenchmarkBubbleSort8(b *testing.B) {
	benchmarkBubbleSort(8, b)
}

func BenchmarkBubbleSort9(b *testing.B) {
	benchmarkBubbleSort(9, b)
}

func BenchmarkBubbleSort10(b *testing.B) {
	benchmarkBubbleSort(10, b)
}

func BenchmarkBubbleSort20(b *testing.B) {
	benchmarkBubbleSort(20, b)
}

func BenchmarkBubbleSort30(b *testing.B) {
	benchmarkBubbleSort(30, b)
}

func BenchmarkBubbleSort40(b *testing.B) {
	benchmarkBubbleSort(40, b)
}

func BenchmarkBubbleSort50(b *testing.B) {
	benchmarkBubbleSort(50, b)
}

func BenchmarkBubbleSort60(b *testing.B) {
	benchmarkBubbleSort(60, b)
}

func BenchmarkBubbleSort70(b *testing.B) {
	benchmarkBubbleSort(70, b)
}

func BenchmarkBubbleSort80(b *testing.B) {
	benchmarkBubbleSort(80, b)
}

func BenchmarkBubbleSort90(b *testing.B) {
	benchmarkBubbleSort(90, b)
}

func BenchmarkBubbleSort100(b *testing.B) {
	benchmarkBubbleSort(100, b)
}

func BenchmarkBubbleSort200(b *testing.B) {
	benchmarkBubbleSort(200, b)
}

func BenchmarkBubbleSort300(b *testing.B) {
	benchmarkBubbleSort(300, b)
}

func BenchmarkBubbleSort400(b *testing.B) {
	benchmarkBubbleSort(400, b)
}

func BenchmarkBubbleSort500(b *testing.B) {
	benchmarkBubbleSort(500, b)
}

func BenchmarkBubbleSort600(b *testing.B) {
	benchmarkBubbleSort(600, b)
}

func BenchmarkBubbleSort700(b *testing.B) {
	benchmarkBubbleSort(700, b)
}

func BenchmarkBubbleSort800(b *testing.B) {
	benchmarkBubbleSort(800, b)
}

func BenchmarkBubbleSort900(b *testing.B) {
	benchmarkBubbleSort(900, b)
}

func BenchmarkBubbleSort1000(b *testing.B) {
	benchmarkBubbleSort(1000, b)
}
