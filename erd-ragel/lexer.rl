package main

%%{
	machine lexer;
	write data;
	access lex.;
	variable p lex.p;
	variable pe lex.pe;
	# Required for nested.
	variable top lex.top;
	variable stack lex.stack;
}%%

type Lexer struct {
	data []byte
	p, pe, cs int
	ts, te, act int
	top int
	stack []int
	result string
}

func NewLexer(data []byte) *Lexer {
	lex := &Lexer{
		data: data,
		pe: len(data),
		stack: make([]int, 0),
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
		#title_name = ('T'|'t')('I'|'i')('T'|'t')('L''l')('E''e');
		title_name = 'title'i; # Case insensitive;
		title = ( alnum | space )+;

		action set_title {
			tok = TITLE;
			fbreak;
		}

		entity := |*
			title_name => { tok = TITLE; fbreak; };
			title => { lval.title = lex.string(); tok = TITLE_VALUE; fbreak; };
			any => { tok = int(lex.data[lex.ts]); fbreak; };
		*|;

		main := ( title_name ':' title >set_title @{ fcall entity; })*;
		write exec;
	}%%

	return tok
}

func (lex *Lexer) string() string {
	return string(lex.data[lex.ts:lex.te])
}
