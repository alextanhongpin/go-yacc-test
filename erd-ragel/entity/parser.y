%{
package entity_parser

func setResult(l yyLexer, val Result) {
	l.(*Lexer).result = val
}
%}

%union {
	str string
	entity Entity
	entities []Entity
	attributes []Attribute
	attribute Attribute
	isPrimary bool
	isForeign bool
}

%token NEWLINE BREAK
%token <str> ENTITY STRING PRIMARY_KEY FOREIGN_KEY ATTRIBUTE

%type <attribute> attribute
%type <entities> main entities
%type <entity> entity
%type <attributes> attributes

%start main

%%

main: entities
		{
			setResult(yylex, $$);
		}

// Entities are separated by newline.
entities: entities BREAK entity
				{
					$$ = append($$, $3);
				}
				| entity
				{
					$$ = []Entity{$1};
				}

entity: '[' STRING ']' NEWLINE attributes
			{
				$$.name = $2;
				$$.attributes = $5;
			}

// Attributes are separated by newline.
attributes: attributes NEWLINE attribute
					{
						$$ = append($$, $3);
					}
					| attribute
					{
						$$ = []Attribute{$1};
					}

attribute: PRIMARY_KEY ATTRIBUTE
				 {
				 	 $$.isPrimary = true;
				 	 $$.field = $2;
				 }
				 | FOREIGN_KEY ATTRIBUTE
				 {
				 	 $$.isForeign = true;
				 	 $$.field = $2;
				 }
				 | ATTRIBUTE
				 {
				 	 $$.field = $1;
		 		 }
