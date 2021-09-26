package main

import "log"

//go:generate goyacc -o parser.go parser.y
//go:generate ragel -Z -G2 -o lexer.go lexer.rl

func main() {
	lexer := NewLexer([]byte(`title: hello world
[Users]
age`))
	log.Println(yyParse(lexer))
	log.Println(lexer.result)
}
