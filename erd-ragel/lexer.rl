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
	relation Relation
}

type Relation struct {
	from string
	to string
	fromCardinality string
	toCardinality string
}

type Entity struct {
	name string
	attributes []string
	relations []Relation
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
	yyErrorVerbose = true

	%%{
		title_name = 'title'i; # Case insensitive;
		char = (alpha | ',' | '(' | ')' | '_' );
		chars = alpha(char|' ')*char;
		quoted_string = '`' (alpha | ',' | '(' | ')' | '_'| ' ')+ '`';
		string = (('*'|'+')?(chars|quoted_string));
		break = [\r\n]{2,};
		newline = [\r\n];
		cardinality = (' '('?'|'*'|'1'|'+')|('?'|'*'|'1'|'+')' ');

		main := |*
			cardinality => { println("cardinality", lex.string()); lval.str = lex.string(); tok = CARDINALITY; fbreak; };
			title_name => { tok = TITLE; fbreak; };
			string => { lval.str = lex.string(); println("string:", lex.string()); tok = STRING; fbreak; };
			break => { println("break"); tok = BREAK; fbreak; };
			# newline => { println("newline", lex.te, eof); if lex.te == eof { fbreak; } else { tok = NEWLINE; } };
			# NOTE: We skip the last EOF if exists ...
			newline => { println("newline", lex.te, eof); if lex.te != eof { tok = NEWLINE }; fbreak; };
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
