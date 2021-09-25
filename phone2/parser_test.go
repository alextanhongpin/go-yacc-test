package phone

import (
	"errors"
	"testing"
)

func TestPhoneParser(t *testing.T) {
	tests := []struct {
		input string
		err   error
	}{
		{"123-456-7890", nil},
		{"1234567890", nil},
		{"123 4567890", errors.New("syntax error")},
	}
	for _, test := range tests {
		res, err := Parse([]byte(test.input))
		if test.err != nil && test.err.Error() != err.Error() {
			t.Fatal(err)
		}
		t.Log(res, err)
	}
}
