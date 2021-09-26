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
	result []Student
}

type Student struct {
	school string
	name string
	modules []string
	group string
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
		main := |*
			[A-Z][a-z\'][\-A-Za-z\', ]* => { lval.name = lex.string(); tok = name; fbreak; };
			[A-Z]+ => { lval.school = lex.string(); tok = school; fbreak; };
			( [A-Z][0-9] | [_]{2} ) => { lval.group = lex.string(); tok = group; fbreak; };
			[A-Z][A-Z][0-9]{3,4} => { lval.module = lex.string(); tok = module; fbreak; };
			[ \t\n];
		*|;

		write exec;
	}%%

	return tok
}

func (lex *Lexer) string() string {
	return string(lex.data[lex.ts:lex.te])
}
