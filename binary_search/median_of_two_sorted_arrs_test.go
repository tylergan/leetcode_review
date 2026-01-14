package binary_search

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
	}{
		{
			name:  "example 1",
			nums1: []int{1, 2},
			nums2: []int{3},
			want:  2.0,
		},
		{
			name:  "example 2",
			nums1: []int{1, 3},
			nums2: []int{2, 4},
			want:  2.5,
		},
		{
			name:  "nums1 empty odd length",
			nums1: []int{},
			nums2: []int{2, 3, 4},
			want:  3.0,
		},
		{
			name:  "nums1 empty even length",
			nums1: []int{},
			nums2: []int{2, 3, 4, 5},
			want:  3.5,
		},
		{
			name:  "nums2 empty odd length",
			nums1: []int{7},
			nums2: []int{},
			want:  7.0,
		},
		{
			name:  "odd total length",
			nums1: []int{1, 2, 3},
			nums2: []int{4, 5},
			want:  3.0,
		},
		{
			name:  "even total length",
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.5,
		},
		{
			name:  "one array longer",
			nums1: []int{1, 2, 3, 4, 5},
			nums2: []int{6, 7},
			want:  4.0,
		},
		{
			name:  "negative values",
			nums1: []int{-5, -3, -1},
			nums2: []int{-2, 0, 2},
			want:  -1.5,
		},
		{
			name:  "all values in nums2 smaller",
			nums1: []int{10, 11, 12},
			nums2: []int{1, 2, 3, 4},
			want:  4.0,
		},
		{
			name:  "all values in nums1 smaller",
			nums1: []int{1, 2, 3, 4},
			nums2: []int{10, 11, 12},
			want:  4.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMedianSortedArrays(tt.nums1, tt.nums2)
			if got != tt.want {
				t.Fatalf("findMedianSortedArrays(%v, %v) = %v, want %v", tt.nums1, tt.nums2, got, tt.want)
			}
		})
	}
}
