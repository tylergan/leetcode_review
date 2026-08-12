package two_pointers

import (
	"math/rand"
	"testing"
)

func TestSortColors(t *testing.T) {
	// Both implementations must produce identical, sorted output, so every
	// case runs against each.
	impls := map[string]func([]int){
		"dutch_flag": sortColors,
		"counting":   sortColorsCounting,
	}

	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "example 1",
			nums: []int{1, 0, 1, 2},
			want: []int{0, 1, 1, 2},
		},
		{
			name: "example 2",
			nums: []int{2, 1, 0},
			want: []int{0, 1, 2},
		},
		{
			name: "single element",
			nums: []int{1},
			want: []int{1},
		},
		{
			name: "already sorted",
			nums: []int{0, 0, 1, 1, 2, 2},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "reverse sorted",
			nums: []int{2, 2, 1, 1, 0, 0},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "all same color",
			nums: []int{1, 1, 1},
			want: []int{1, 1, 1},
		},
		{
			name: "no whites",
			nums: []int{2, 0, 2, 0},
			want: []int{0, 0, 2, 2},
		},
		{
			name: "trailing two triggers re-check",
			nums: []int{0, 2, 0},
			want: []int{0, 0, 2},
		},
	}

	for implName, impl := range impls {
		for _, tt := range tests {
			t.Run(implName+"/"+tt.name, func(t *testing.T) {
				nums := append([]int(nil), tt.nums...)
				impl(nums)
				for i := range tt.want {
					if nums[i] != tt.want[i] {
						t.Fatalf("%s(%v) = %v, want %v", implName, tt.nums, nums, tt.want)
					}
				}
			})
		}
	}
}

// Random inputs cross-checked by counting, with both impls expected to agree.
func TestSortColorsRandom(t *testing.T) {
	rng := rand.New(rand.NewSource(2024))

	for iter := 0; iter < 1000; iter++ {
		n := rng.Intn(300) + 1
		nums := make([]int, n)
		var freq [3]int
		for i := range nums {
			nums[i] = rng.Intn(3)
			freq[nums[i]]++
		}

		want := make([]int, 0, n)
		for c := 0; c < 3; c++ {
			for j := 0; j < freq[c]; j++ {
				want = append(want, c)
			}
		}

		for implName, impl := range map[string]func([]int){"dutch_flag": sortColors, "counting": sortColorsCounting} {
			got := append([]int(nil), nums...)
			impl(got)
			for i := range want {
				if got[i] != want[i] {
					t.Fatalf("iter %d %s: got %v, want %v (input %v)", iter, implName, got, want, nums)
				}
			}
		}
	}
}
