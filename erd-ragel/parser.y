%{
package main

func setResult(l yyLexer, val Erd) {
	l.(*Lexer).result = val
}
%}

%union {
	str string
	title string
	attributes []string
	attribute string
	entity Entity
	entities []Entity
	result Erd
}

%token LBRAC RBRAC TITLE STRING NEWLINE BREAK


%type <result> main
%type <str> STRING
%type <title> title
%type <entity> entity
%type <entities> entities
%type <attributes> attributes
%type <attribute> attribute

%start main

%%

main: title
		{
			$$.title = $1;
			setResult(yylex, $$);
		}
		| title BREAK entities
		{
			$$.title = $1;
			$$.entities = $3;
			setResult(yylex, $$);
		}
		;

title: TITLE ':' STRING { $$ = $3; }
		 ;

// Entities are separated by newline.
entities: entity { $$ = []Entity{$1}; }
				| entity BREAK entities { $$ = append($$, $1); }
				;

entity: LBRAC STRING RBRAC NEWLINE attributes { $$.name = $2; $$.attributes = $5; }
			;

// Attributes are separated by newline.
attributes: attribute { $$ = []string{$1}; }
					| attribute NEWLINE attributes { $$ = append($$, $1); }
					;

attribute: STRING { $$ = $1; }
				 ;
