package jsonparser

import "testing"

func TestParse(t *testing.T) {
	res, err := Parse([]byte(`{}`))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
