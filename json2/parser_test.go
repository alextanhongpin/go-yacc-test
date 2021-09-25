package jsonparser

import "testing"

func TestParse(t *testing.T) {
	result, err := Parse([]byte(`{"age": 1, "name": "john", "isMarried": false}`))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
