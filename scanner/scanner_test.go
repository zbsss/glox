//go:build unit
// +build unit

package scanner

import (
	"testing"
)

func TestScanTwoCharacterTokens(t *testing.T) {
	tests := []struct {
		source string
		want   []string
	}{
		{source: "!==", want: []string{"!=", "="}},
		{source: "===", want: []string{"==", "="}},
		{source: "<=>=", want: []string{"<=", ">="}},
	}

	for _, tc := range tests {
		sc := NewScanner(tc.source)

		tokens := sc.ScanTokens()

		for i, token := range tokens[:len(tokens)-1] {
			if tc.want[i] != token.Lexeme {
				t.Errorf("ScanTokens(%v): expected %v token, but got %v instead", tc.source, tc.want[i], token.Lexeme)
			}
		}
	}
}
