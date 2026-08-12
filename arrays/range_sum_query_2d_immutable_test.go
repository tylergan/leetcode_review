package arrays

import (
	"math/rand"
	"testing"
)

func TestNumMatrixExample(t *testing.T) {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	nm := Constructor(matrix)

	queries := []struct {
		row1, col1, row2, col2 int
		want                   int
	}{
		{2, 1, 4, 3, 8},
		{1, 1, 2, 2, 11},
		{1, 2, 2, 4, 12},
	}

	for _, q := range queries {
		if got := nm.SumRegion(q.row1, q.col1, q.row2, q.col2); got != q.want {
			t.Fatalf("SumRegion(%d,%d,%d,%d) = %d, want %d", q.row1, q.col1, q.row2, q.col2, got, q.want)
		}
	}
}

func TestNumMatrixSingleCell(t *testing.T) {
	nm := Constructor([][]int{{42}})
	if got := nm.SumRegion(0, 0, 0, 0); got != 42 {
		t.Fatalf("SumRegion(0,0,0,0) = %d, want 42", got)
	}
}

func TestNumMatrixSingleCellRegions(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	nm := Constructor(matrix)

	for i := range matrix {
		for j := range matrix[i] {
			if got := nm.SumRegion(i, j, i, j); got != matrix[i][j] {
				t.Fatalf("SumRegion(%d,%d,%d,%d) = %d, want %d", i, j, i, j, got, matrix[i][j])
			}
		}
	}
}

func TestNumMatrixNegatives(t *testing.T) {
	matrix := [][]int{
		{-1, -2},
		{-3, -4},
	}
	nm := Constructor(matrix)

	if got := nm.SumRegion(0, 0, 1, 1); got != -10 {
		t.Fatalf("full region = %d, want -10", got)
	}
	if got := nm.SumRegion(1, 0, 1, 1); got != -7 {
		t.Fatalf("bottom row = %d, want -7", got)
	}
}

// bruteRegion sums a rectangle directly for cross-checking.
func bruteRegion(matrix [][]int, row1, col1, row2, col2 int) int {
	sum := 0
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			sum += matrix[i][j]
		}
	}
	return sum
}

func TestNumMatrixMatchesBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))

	for iter := 0; iter < 200; iter++ {
		m := rng.Intn(20) + 1
		n := rng.Intn(20) + 1
		matrix := make([][]int, m)
		for i := range matrix {
			matrix[i] = make([]int, n)
			for j := range matrix[i] {
				matrix[i][j] = rng.Intn(20001) - 10000 // [-10000, 10000]
			}
		}

		nm := Constructor(matrix)

		for q := 0; q < 50; q++ {
			row1 := rng.Intn(m)
			row2 := row1 + rng.Intn(m-row1)
			col1 := rng.Intn(n)
			col2 := col1 + rng.Intn(n-col1)

			got := nm.SumRegion(row1, col1, row2, col2)
			want := bruteRegion(matrix, row1, col1, row2, col2)
			if got != want {
				t.Fatalf("iter %d: SumRegion(%d,%d,%d,%d) = %d, want %d", iter, row1, col1, row2, col2, got, want)
			}
		}
	}
}
