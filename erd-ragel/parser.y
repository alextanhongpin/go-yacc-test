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
	relation Relation
	result Erd
}

%token LBRAC RBRAC TITLE CARDINALITY STRING QUOTED_STRING NEWLINE BREAK

%type <result> main
%type <str> STRING CARDINALITY
%type <title> title
%type <entity> entity
%type <entities> entities
%type <attributes> attributes
%type <attribute> attribute
%type <relation> relation

%start main

%%

main: title
		{
			$$.title = $1;
			setResult(yylex, $$);
		}
		| title BREAK entities BREAK relation
		{
			$$.title = $1;
			$$.entities = $3;
			$$.relation = $5;
			setResult(yylex, $$);
		}

// A title is separated by colon.
title: TITLE ':' STRING { $$ = $3; }

relation: STRING CARDINALITY '-' '-' CARDINALITY STRING { $$.from = $1; $$.fromCardinality = $2; $$.toCardinality = $5; $$.to = $6; }

// Entities are separated by newline.
entities: entities BREAK entity { $$ = append($$, $3); }
				| entity { $$ = []Entity{$1}; }

entity: '[' STRING ']' NEWLINE attributes { $$.name = $2; $$.attributes = $5; }

// Attributes are separated by newline.
attributes: attribute { $$ = []string{$1}; }
					| attribute NEWLINE attributes { $$ = append($$, $1); }

attribute: STRING { $$ = $1; }
