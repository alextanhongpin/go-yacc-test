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

type Result []Entity

type Lexer struct {
	data []byte
	p, pe, cs int
	ts, te, act int
	stack []int
	top int
	result []Entity
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
		word = (alpha | '_' );
		string = (word+|'`' word(word|' ')*word+ '`'| "'" word(word|' ')*word "'" | '"' word(word|' ')*word '"');
		break = [\r\n]{2,};
		nl = [\r\n];
		attribute = (word | '-' | '(' | ')' | ' ' | ',')+;

		prepush {
			// Dynamically resize the stack.
			for len(lex.stack) < lex.top + 1 {
				lex.stack = append(lex.stack, 0)
			}
		}

		entity := |*
			string    { tok = STRING; lval.str = lex.string(); fbreak; };
			'[' | ']' { tok = lex.token(); fbreak; };
			nl 				{ fhold; fret; }; # Return to main without consuming token.
		*|;

		main := |*
			'*' => { lval.str = lex.string(); println("*:", lex.string()); tok = PRIMARY_KEY; fbreak; };
			'+' => { lval.str = lex.string(); println("+:", lex.string()); tok = FOREIGN_KEY; fbreak; };
			string => { lval.str = lex.string(); println("string:", lex.string()); tok = STRING; fbreak; };
			attribute => { lval.str = lex.string(); tok = ATTRIBUTE; fbreak; };
			break => { tok = BREAK; fbreak; };
			# nl => { println("nl", lex.te, eof); if lex.te == eof { fbreak; } else { tok = NEWLINE; } };
			# NOTE: We skip the last EOF if exists ...
			nl => { if lex.te != eof { tok = NEWLINE }; fbreak; };
			'[' { fhold; fcall entity; }; # On detecting the first bracket, handle the substate.
			' ';
			any => { tok = lex.token(); fbreak; };
		*|;

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
