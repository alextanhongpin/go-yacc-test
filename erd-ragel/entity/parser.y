%{
package entity_parser

func setResult(l yyLexer, val Result) {
	l.(*Lexer).result = val
}
%}

%union {
	str string
	isPrimary bool
	isForeign bool
	attribute Attribute
	attributes []Attribute
	entity Entity
	relation Relation

	result Result
}

%token <str> CARDINALITY ENTITY ATTRIBUTE PRIMARY_KEY FOREIGN_KEY

%type <attribute> attribute
%type <attributes> attributes
%type <entity> entity
%type <relation> relation
%type <result> main body

%start main

%%

// Body can be entity or relation, in any order.
body: entity
		{
			$$.entities = []Entity{$1};
		}
		| relation
		{
			$$.relations = []Relation{$1};
		}

// Main can have one or more body.
main: body
		{
			$$ = $1;
		}
		| body main
		{
			$$.relations = append($$.relations, $2.relations...);
			$$.entities = append($$.entities, $2.entities...);
			setResult(yylex, $$);
		}

entity: ENTITY attributes
			{
				$$.name = $1;
				$$.attributes = $2;
			}

attributes: attributes attribute
					{
						$$ = append($$, $2);
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

relation: ENTITY CARDINALITY CARDINALITY ENTITY
				{
					$$.from = $1;
					$$.fromCardinality = $3; // Note that the cardinality is the opposite.
					$$.toCardinality = $2;
					$$.to = $4;
				}
