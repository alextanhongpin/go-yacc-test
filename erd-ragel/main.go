package main

import (
	_ "embed"
	"log"
)

//go:generate goyacc -o parser.go parser.y
//go:generate ragel -Z -G2 -o lexer.go lexer.rl
// go:generate ragel -V lexer.rl -o lexer.dot
// go:generate dot -Tpng lexer.dot > lexer.png

//go:embed input.txt
var input []byte

func main() {
	lexer := NewLexer(input)
	log.Println(yyParse(lexer))
	log.Printf("%+v\n", lexer.result)
}
