%{
package main

func appendResult(l yyLexer, val string) {
	l.(*Lexer).result = append(l.(*Lexer).result, val)
}
%}

%union {
	phone string
	token string
}

%token <token> COUNTRY OPERATOR PHONE

%type <phone> phone

%%

start: phone
		 {
		 		appendResult(yylex, $1)
		 }
		 | start ',' phone
		 {
		 		appendResult(yylex, $3)
		 }

phone: '+' COUNTRY '(' OPERATOR ')' PHONE
		 {
		 	$$ = "+" + $2 + "(" + $4 + ")" + $6
	   }
		 | COUNTRY OPERATOR PHONE
		 { $$ = "+" + $1 + "(" + $2 + ")" + $3 }
