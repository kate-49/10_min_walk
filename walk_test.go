package walk

import (
	"testing"
)

func TestRun(t *testing.T) {

	tests := []struct {
		input []string
		want  bool
	}{
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
