package main

%%{
	machine lexer;
	write data;
	access lex.;
	variable p lex.p;
	variable pe lex.pe;
}%%

type Lexer struct {
	data []byte
	p, pe, cs int
	ts, te, act int
	result int
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
		title = ( alnum | space )+;

		main := |*
			title_name => { tok = TITLE; fbreak; };
			title => { tok = TITLE_VALUE; fbreak; };
			alnum+ => { tok = FIELD; fbreak; };
			space+ => { tok = NEWLINE; fbreak; };
			'['alnum+']'=> { tok = ENTITY; fbreak; };
			any => { tok = int(lex.data[lex.ts]); fbreak; };
		*|;

		write exec;
	}%%

	return tok
}

func (lex *Lexer) string() string {
	return string(lex.data[lex.ts:lex.te])
}
