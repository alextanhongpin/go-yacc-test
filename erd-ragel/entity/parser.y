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
	entities []Entity
	entity Entity
	relations []Relation
	relation Relation

	result Result
}

%token <str> CARDINALITY ENTITY ATTRIBUTE PRIMARY_KEY FOREIGN_KEY

%type <attribute> attribute
%type <attributes> attributes
%type <entities> entities
%type <entity> entity
%type <relation> relation
%type <relations> relations
%type <result> main

%start main

%%

main: entities
		{
			$$.entities = $1;
			setResult(yylex, $$);
		}
		| relations
		{
			$$.relations = $1;
			setResult(yylex, $$);
		}
		| entities relations
		{
			$$.entities = $1;
			$$.relations = $2;
			setResult(yylex, $$);
		}
		| relations entities
		{
			$$.relations = $1;
			$$.entities = $2;
			setResult(yylex, $$);
		}

entities: entities entity
				{
					$$ = append($$, $2);
				}
				| entity
				{
					$$ = []Entity{$1};
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

relations: relations relation
				 {
				   $$ = append($$, $2);
				 }
				 | relation
				 {
					 $$ = []Relation{$1};
				 }

relation: ENTITY CARDINALITY CARDINALITY ENTITY
			{
				$$.from = $1;
				$$.fromCardinality = $3; // Note that the cardinality is the opposite.
				$$.toCardinality = $2;
				$$.to = $4;
			}
