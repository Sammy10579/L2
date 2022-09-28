package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestUnpack(t *testing.T) {
	unpackTests := []struct {
		arg         string
		expected    string
		expectedErr error
	}{
		{`a4bc2d5e`, "aaaabccddddde", nil},
		{`abcd`, "abcd", nil},
		{`45`, "", fmt.Errorf("incorrent string")},
		{"", "", nil},
		{`qwe\4\5`, "qwe45", nil},
		{`qwe\45`, "qwe44444", nil},
		{`qwe\\5`, `qwe\\\\\`, nil},
	}

	for _, test := range unpackTests {
		output, err := Unpack(test.arg)
		if output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
		if err == nil && test.expectedErr == nil {
			if errors.Is(err, test.expectedErr) != true {
				t.Errorf("Output %q not equal to expected %q", err, test.expectedErr)
			}
		}
	}
}
