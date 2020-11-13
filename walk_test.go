package walk

import (
	"testing"
)

func TestRun(t *testing.T) {

	tests := []struct {
		input []string
		want  bool
	}{
		//test length less than 10 fails
		{input: []string{"w", "s", "e", "s", "w", "w"}, want: false},
		//test length more than 10 fails
		{input: []string{"w", "s", "e", "s", "w", "w", "n", "n", "e", "s", "w", "w"}, want: false},
		// acceptance criteria tests
		{input: []string{"w", "s", "e", "e", "n", "n", "e", "s", "w", "w"}, want: true},
	}

	for _, tc := range tests {
		got := Run(tc.input)
		if tc.want != got {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
