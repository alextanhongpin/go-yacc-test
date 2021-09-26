%{
package main

func setResult(l yyLexer, val int) {
	l.(*Lexer).result = val
}
%}

%union {
	result int
}

%token TITLE TITLE_VALUE FIELD ENTITY NEWLINE

%start main

%%

main: title NEWLINE entities
		{
			setResult(yylex, 0);
		}
		;


title: TITLE ':' TITLE_VALUE
		 ;

entities: entity
				| entity entities
				;

entity: '[' ENTITY ']' attributes
			;

attributes:
					| attributes FIELD;
