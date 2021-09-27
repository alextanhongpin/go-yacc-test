package entity_parser

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input []byte

func TestParser(t *testing.T) {
	Parse(input)
}
