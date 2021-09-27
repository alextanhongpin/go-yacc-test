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
		quoted_string = '`' word+ '`';
		string = (word+|'`' word+ '`'| '' word+ '' | " word+ ");
		break = [\r\n]{2,};
		newline = [\r\n];
		attribute = (word | '-' | '(' | ')' | ' ' | ',')+;

		main := |*
			'*' => { lval.str = lex.string(); println("*:", lex.string()); tok = PRIMARY_KEY; fbreak; };
			'+' => { lval.str = lex.string(); println("+:", lex.string()); tok = FOREIGN_KEY; fbreak; };
			string => { lval.str = lex.string(); println("string:", lex.string()); tok = STRING; fbreak; };
			attribute => { lval.str = lex.string(); println("attribute:", lex.string()); tok = ATTRIBUTE; fbreak; };
			break => { println("break"); tok = BREAK; fbreak; };
			# newline => { println("newline", lex.te, eof); if lex.te == eof { fbreak; } else { tok = NEWLINE; } };
			# NOTE: We skip the last EOF if exists ...
			newline => { println("newline"); if lex.te != eof { tok = NEWLINE }; fbreak; };
			' ';
			any => { println(string(lex.data[lex.ts])); tok = int(lex.data[lex.ts]); fbreak; };
		*|;

		write exec;

	}%%

	return tok
}

func (lex *Lexer) string() string {
	return string(lex.data[lex.ts:lex.te])
}
