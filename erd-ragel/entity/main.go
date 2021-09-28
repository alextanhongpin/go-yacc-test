package entity_parser

import (
	_ "embed"
	"log"
)

//go:generate goyacc -o parser.go parser.y
//go:generate ragel -Z -G2 -o lexer.go lexer.rl
// go:generate ragel -V lexer.rl -o lexer.dot
// go:generate dot -Tpng lexer.dot > lexer.png

func Parse(input []byte) {
	lexer := NewLexer(input)
	log.Println(yyParse(lexer))
	for i, ent := range lexer.result.entities {
		log.Printf("%d. %+v\n", i+1, ent.name)
		for j, attr := range ent.attributes {
			var symbol string
			if attr.isPrimary {
				symbol = "*"
			} else if attr.isForeign {
				symbol = "+"
			}
			log.Printf("%d. %s%+v\n", j+1, symbol, attr.field)
		}
		log.Println("")
	}

	for i, relation := range lexer.result.relations {
		log.Printf("%d. %+v\n", i+1, relation)
	}
}
