package two_pointers

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "example 1: palindrome with spaces and punctuation",
			s:    "Was it a car or a cat I saw?",
			want: true,
		},
		{
			name: "example 2: not a palindrome",
			s:    "tab a cat",
			want: false,
		},
		{
			name: "single character",
			s:    "a",
			want: true,
		},
		{
			name: "single alphanumeric",
			s:    "1",
			want: true,
		},
		{
			name: "simple palindrome lowercase",
			s:    "racecar",
			want: true,
		},
		{
			name: "simple palindrome mixed case",
			s:    "RaceCar",
			want: true,
		},
		{
			name: "palindrome with spaces",
			s:    "race car",
			want: true,
		},
		{
			name: "not palindrome simple",
			s:    "hello",
			want: false,
		},
		{
			name: "empty string after removing non-alphanumeric",
			s:    " ",
			want: true,
		},
		{
			name: "only special characters",
			s:    "!!!",
			want: true,
		},
		{
			name: "palindrome with numbers",
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "palindrome numbers only",
			s:    "12321",
			want: true,
		},
		{
			name: "not palindrome numbers",
			s:    "12345",
			want: false,
		},
		{
			name: "palindrome with mixed alphanumeric",
			s:    "0P",
			want: false,
		},
		{
			name: "palindrome alphanumeric mixed",
			s:    "a1b2b1a",
			want: true,
		},
		{
			name: "not palindrome alphanumeric",
			s:    "a1b2c1a",
			want: false,
		},
		{
			name: "palindrome with lots of punctuation",
			s:    "A man, a plan, a canal: Panama!",
			want: true,
		},
		{
			name: "single word palindrome",
			s:    "noon",
			want: true,
		},
		{
			name: "single word not palindrome",
			s:    "world",
			want: false,
		},
		{
			name: "palindrome all caps",
			s:    "RACECAR",
			want: true,
		},
		{
			name: "two characters palindrome",
			s:    "aa",
			want: true,
		},
		{
			name: "two characters not palindrome",
			s:    "ab",
			want: false,
		},
		{
			name: "palindrome with underscores and dashes",
			s:    "race_-_car",
			want: true,
		},
		{
			name: "complex punctuation palindrome",
			s:    "A Santa at NASA",
			want: true,
		},
		{
			name: "even length palindrome",
			s:    "abba",
			want: true,
		},
		{
			name: "odd length palindrome",
			s:    "aba",
			want: true,
		},
		{
			name: "even length not palindrome",
			s:    "abcd",
			want: false,
		},
		{
			name: "palindrome with comma",
			s:    "a,b,a",
			want: true,
		},
		{
			name: "palindrome numbers and letters",
			s:    "1a2a1",
			want: true,
		},
		{
			name: "all uppercase with spaces",
			s:    "NEVER ODD OR EVEN",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.s)
			if got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
