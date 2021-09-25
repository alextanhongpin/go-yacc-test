%{
package jsonparser

type pair struct {
	key string
	val interface{}
}

func setResult(l yyLexer, v map[string]interface{}) {
	// yyLexer is an interface, cast them
	// back to it's original struct pointer
	// and set the value.
	l.(*lex).result = v
}
%}


// Union allows us to set additional fields to yyval.
%union {
	obj map[string]interface{}
	list []interface{}
	pair pair
	val interface{}
}

%token LexError
%token <val> String Number Literal

%type <obj> object members
%type <pair> pair
%type <val> array
%type <list> elements
%type <val> value

%start object

%%

object: '{' members '}' { $$ = $2; setResult(yylex, $$); }

members: /* empty */{ $$ = map[string]interface{}{} }
			 | pair { $$ = map[string]interface{}{$1.key: $1.val} }
			 | members ',' pair { $1[$3.key] = $3.val; $$ = $1; }

pair: String ':' value { $$ = pair{key: $1.(string), val: $3} }

array: '[' elements ']' { $$ = $2 }
elements: { $$ = []interface{}{} }
				| value { $$ = []interface{}{$1} }
				| elements ',' value { $$ = append($1, $3) }


value: String
		 | Number
		 | Literal
		 | object { $$ = $1 }
		 | array
