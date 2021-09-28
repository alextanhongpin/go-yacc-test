package entity_parser

%%{
	machine lexer;
	write data;
	access lex.;
	variable p lex.p;
	variable pe lex.pe;
}%%


type Attribute struct {
	isPrimary bool
	isForeign bool
	field string
}

type Entity struct {
	name string
	attributes []Attribute
}

type Relation struct {
	from string
	to string
	fromCardinality string
	toCardinality string
}

type Result struct {
	entities []Entity
	relations []Relation
}

type Lexer struct {
	data []byte
	p, pe, cs int
	ts, te, act int
	stack []int
	top int

	result Result
}

func NewLexer(data []byte) *Lexer {
	lex := &Lexer{
		data: data,
		pe: len(data),
	}
	%% write init;
	return lex
}

func (l *Lexer) Error(msg string) {
	println(msg)
}

func (lex *Lexer) Lex(lval *yySymType) int {
	eof := lex.pe
	var tok int
	yyErrorVerbose = true

	%%{
		ch = (alpha | '_' );
		string = (ch+|'`' ch(ch|' ')*ch+ '`'| "'" ch(ch|' ')*ch "'" | '"' ch(ch|' ')*ch '"');
		attribute = ch(ch | '-' | '(' | ')' | ' ' | ',')*;
		cardinality = ('?' | '1' | '*' | '+');
		newline = [\r\n];

		prepush {
			// Dynamically resize the stack.
			for len(lex.stack) < lex.top + 1 {
				lex.stack = append(lex.stack, 0)
			}
		}

		entity := |*
			# Trims the first and last character.
			'[' string ']'{ tok = ENTITY; lval.str = string(lex.data[lex.ts+1:lex.te-1]); fbreak; };
			'*' 					{ tok = PRIMARY_KEY; lval.str = lex.string(); fbreak; };
			'+' 					{ tok = FOREIGN_KEY; lval.str = lex.string(); fbreak; };
			attribute 		{ tok = ATTRIBUTE; lval.str = lex.string(); fbreak; };
			newline{2,} 	{ fret; };
			newline;
		*|;

		relation := |*
			cardinality { tok = CARDINALITY; lval.str = lex.string(); fbreak; };
			string 			{ tok = ENTITY; lval.str = lex.string(); fbreak; };
			newline{2,} { fret; }; # Return to the main machine once two or more newline is found.
			('-'|' '); # Skips dashes and spaces.
			newline;   # Skips newline.
		*|;

		action call_entity {
			// fhold does not consume the token.
			fhold; fcall entity;
		}

		action call_relation {
			// fhold does not consume the token.
			fhold; fcall relation;
		}

		# Body consists of relation or entity.
		body =
			^'[' @call_relation | # Anything that doesn't start with a left bracket is potentially a relation.
			'[' @call_entity;     # Otherwise, it's an entity.

		# Main can have multiple bodies.
		main := body+;

		write exec;

	}%%

	return tok
}

func (lex *Lexer) string() string {
	return string(lex.data[lex.ts:lex.te])
}

func (lex *Lexer) token() int {
	return int(lex.data[lex.ts])
}
