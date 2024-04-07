package strings

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"hello", "olleh"},
		{"world", "dlrow"},
		// Add more test cases
	}

	for _, tc := range tests {
		t.Run("Reverse", func(t *testing.T) {
			got := Reverse(tc.input)
			if got != tc.want {
				t.Errorf("Reverse(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

// TestRuneCount is a test function for the RuneCount function
func TestRuneCount(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"hello", 5},
		{"世界", 2},
		// Add more test cases
	}

	for _, tc := range tests {
		t.Run("RuneCount", func(t *testing.T) {
			got := RuneCount(tc.input)
			if got != tc.want {
				t.Errorf("RuneCount(%q) = %d, want %d", tc.input, got, tc.want)
			}
		})
	}
}
