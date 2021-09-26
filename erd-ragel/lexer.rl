package main

%%{
	machine lexer;
	write data;
	access lex.;
	variable p lex.p;
	variable pe lex.pe;
}%%

type Erd struct {
	title string
	entities []Entity
}

type Entity struct {
	name string
	attributes []string
}

type Lexer struct {
	data []byte
	p, pe, cs int
	ts, te, act int
	result Erd
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

	%%{
		title_name = 'title'i; # Case insensitive;
		string = ('*'|'+')?( alnum | ' ' | ',' | '(' | ')' )+;
		break = [\r\n]{2,};
		newline = [\r\n];

		main := |*
			title_name => { tok = TITLE; fbreak; };
			string => { lval.str = lex.string(); println("string:", lex.string()); tok = STRING; };
			break => { println("break"); tok = BREAK; };
			newline => { println("newline", lex.te, eof); if lex.te == eof { fbreak; } else { tok = NEWLINE; } };
			'[' => { tok = LBRAC; };
			']' => { tok = RBRAC; };
			any => { println(string(lex.data[lex.ts])); tok = int(lex.data[lex.ts]); fbreak; };
		*|;

		write exec;

	}%%

	return tok
}

func (lex *Lexer) string() string {
	return string(lex.data[lex.ts:lex.te])
}
