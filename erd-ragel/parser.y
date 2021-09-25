%{
package main

func setResult(l yyLexer, val string) {
	l.(*Lexer).result = val
}
%}

%union {
	title string
	entity string
	erd string
}

%token <title> TITLE TITLE_VALUE
%token <entity> ENTITY

%type <erd> erd
%type <entity> entity

%start main

%%

main:
	erd
	{
		setResult(yylex, $1)
	}

erd:
	TITLE ':' TITLE_VALUE
	{
		$$ = $3;
	}
| TITLE ':' TITLE_VALUE entity
	{
		$$ = $3;
	}

entity:
	'[' ENTITY ']'
	{
		$$ = $2;
	}
