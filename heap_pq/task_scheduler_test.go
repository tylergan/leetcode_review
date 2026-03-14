package heappq

import "testing"

func TestLeastInterval(t *testing.T) {
	tests := []struct {
		name  string
		tasks []byte
		n     int
		want  int
	}{
		{
			name:  "example 1",
			tasks: []byte{'X', 'X', 'Y', 'Y'},
			n:     2,
			want:  5,
		},
		{
			name:  "example 2",
			tasks: []byte{'A', 'A', 'A', 'B', 'C'},
			n:     3,
			want:  9,
		},
		{
			name:  "no cooldown",
			tasks: []byte{'A', 'A', 'A', 'B', 'B'},
			n:     0,
			want:  5,
		},
		{
			name:  "single task",
			tasks: []byte{'A'},
			n:     5,
			want:  1,
		},
		{
			name:  "all same task",
			tasks: []byte{'A', 'A', 'A'},
			n:     2,
			want:  7,
		},
		{
			name:  "all distinct tasks",
			tasks: []byte{'A', 'B', 'C', 'D'},
			n:     3,
			want:  4,
		},
		{
			name:  "cooldown filled by other tasks",
			tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:     1,
			want:  6,
		},
		{
			name:  "multiple tasks same max freq",
			tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C'},
			n:     2,
			want:  8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastInterval(tt.tasks, tt.n); got != tt.want {
				t.Fatalf("leastInterval(%q, %d) = %d, want %d", tt.tasks, tt.n, got, tt.want)
			}
		})
	}
}
