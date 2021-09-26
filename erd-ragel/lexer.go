
//line lexer.rl:1
package main


//line lexer.go:7
const lexer_start int = 0
const lexer_first_final int = 0
const lexer_error int = -1

const lexer_en_main int = 0


//line lexer.rl:9


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
	
//line lexer.go:41
	{
	 lex.cs = lexer_start
	 lex.ts = 0
	 lex.te = 0
	 lex.act = 0
	}

//line lexer.rl:34
	return lex
}

func (l *Lexer) Error(msg string) {
	println(msg)
}

func (lex *Lexer) Lex(lval *yySymType) int {
	eof := lex.pe
	var tok int

	
//line lexer.go:62
	{
	if ( lex.p) == ( lex.pe) {
		goto _test_eof
	}
	switch  lex.cs {
	case 0:
		goto st_case_0
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	}
	goto st_out
tr0:
//line lexer.rl:60
 lex.te = ( lex.p)+1
{ println(string(lex.data[lex.ts])); tok = int(lex.data[lex.ts]); {( lex.p)++;  lex.cs = 0; goto _out } }
	goto st0
tr5:
//line lexer.rl:58
 lex.te = ( lex.p)+1
{ tok = LBRAC; {( lex.p)++;  lex.cs = 0; goto _out } }
	goto st0
tr6:
//line lexer.rl:59
 lex.te = ( lex.p)+1
{ tok = RBRAC; {( lex.p)++;  lex.cs = 0; goto _out } }
	goto st0
tr7:
//line NONE:1
	switch  lex.act {
	case 1:
	{( lex.p) = ( lex.te) - 1
 tok = TITLE; {( lex.p)++;  lex.cs = 0; goto _out } }
	case 2:
	{( lex.p) = ( lex.te) - 1
 lval.str = lex.string(); println("string:", lex.string()); tok = STRING; {( lex.p)++;  lex.cs = 0; goto _out } }
	case 3:
	{( lex.p) = ( lex.te) - 1
 println("break"); tok = BREAK; {( lex.p)++;  lex.cs = 0; goto _out } }
	case 4:
	{( lex.p) = ( lex.te) - 1
 println("newline", lex.te, eof); if lex.te != eof { tok = NEWLINE }; {( lex.p)++;  lex.cs = 0; goto _out } }
	case 7:
	{( lex.p) = ( lex.te) - 1
 println(string(lex.data[lex.ts])); tok = int(lex.data[lex.ts]); {( lex.p)++;  lex.cs = 0; goto _out } }
	}
	
	goto st0
tr9:
//line lexer.rl:53
 lex.te = ( lex.p)
( lex.p)--
{ lval.str = lex.string(); println("string:", lex.string()); tok = STRING; {( lex.p)++;  lex.cs = 0; goto _out } }
	goto st0
	st0:
//line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof0
		}
	st_case_0:
//line NONE:1
 lex.ts = ( lex.p)

//line lexer.go:137
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr1
		case 13:
			goto tr1
		case 32:
			goto tr2
		case 44:
			goto tr2
		case 84:
			goto st3
		case 91:
			goto tr5
		case 93:
			goto tr6
		case 116:
			goto st3
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			switch {
			case  lex.data[( lex.p)] > 41:
				if 42 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 43 {
					goto tr3
				}
			case  lex.data[( lex.p)] >= 40:
				goto tr2
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr2
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr2
			}
		default:
			goto tr2
		}
		goto tr0
tr1:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:57
 lex.act = 4;
	goto st1
tr8:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:54
 lex.act = 3;
	goto st1
	st1:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof1
		}
	st_case_1:
//line lexer.go:198
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr8
		case 13:
			goto tr8
		}
		goto tr7
tr2:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:53
 lex.act = 2;
	goto st2
tr3:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:60
 lex.act = 7;
	goto st2
tr13:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:52
 lex.act = 1;
	goto st2
	st2:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof2
		}
	st_case_2:
//line lexer.go:232
		switch  lex.data[( lex.p)] {
		case 32:
			goto tr2
		case 44:
			goto tr2
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 41 {
				goto tr2
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr2
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr2
			}
		default:
			goto tr2
		}
		goto tr7
	st3:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof3
		}
	st_case_3:
		switch  lex.data[( lex.p)] {
		case 32:
			goto tr2
		case 44:
			goto tr2
		case 73:
			goto st4
		case 105:
			goto st4
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 41 {
				goto tr2
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr2
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr2
			}
		default:
			goto tr2
		}
		goto tr9
	st4:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof4
		}
	st_case_4:
		switch  lex.data[( lex.p)] {
		case 32:
			goto tr2
		case 44:
			goto tr2
		case 84:
			goto st5
		case 116:
			goto st5
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 41 {
				goto tr2
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr2
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr2
			}
		default:
			goto tr2
		}
		goto tr9
	st5:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof5
		}
	st_case_5:
		switch  lex.data[( lex.p)] {
		case 32:
			goto tr2
		case 44:
			goto tr2
		case 76:
			goto st6
		case 108:
			goto st6
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 41 {
				goto tr2
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr2
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr2
			}
		default:
			goto tr2
		}
		goto tr9
	st6:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof6
		}
	st_case_6:
		switch  lex.data[( lex.p)] {
		case 32:
			goto tr2
		case 44:
			goto tr2
		case 69:
			goto tr13
		case 101:
			goto tr13
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 41 {
				goto tr2
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr2
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr2
			}
		default:
			goto tr2
		}
		goto tr9
	st_out:
	_test_eof0:  lex.cs = 0; goto _test_eof
	_test_eof1:  lex.cs = 1; goto _test_eof
	_test_eof2:  lex.cs = 2; goto _test_eof
	_test_eof3:  lex.cs = 3; goto _test_eof
	_test_eof4:  lex.cs = 4; goto _test_eof
	_test_eof5:  lex.cs = 5; goto _test_eof
	_test_eof6:  lex.cs = 6; goto _test_eof

	_test_eof: {}
	if ( lex.p) == eof {
		switch  lex.cs {
		case 1:
			goto tr7
		case 2:
			goto tr7
		case 3:
			goto tr9
		case 4:
			goto tr9
		case 5:
			goto tr9
		case 6:
			goto tr9
		}
	}

	_out: {}
	}

//line lexer.rl:65


	return tok
}

func (lex *Lexer) string() string {
	return string(lex.data[lex.ts:lex.te])
}
