
//line lexer.rl:1
package main


//line lexer.go:7
const lexer_start int = 1
const lexer_first_final int = 1
const lexer_error int = -1

const lexer_en_main int = 1


//line lexer.rl:9


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
	
//line lexer.go:31
	{
	 lex.cs = lexer_start
	 lex.ts = 0
	 lex.te = 0
	 lex.act = 0
	}

//line lexer.rl:24
	return lex
}

func (l *Lexer) Error(msg string) {
	println(msg)
}

func (lex *Lexer) Lex(lval *yySymType) int {
	eof := lex.pe
	var tok int

	
//line lexer.go:52
	{
	if ( lex.p) == ( lex.pe) {
		goto _test_eof
	}
	switch  lex.cs {
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
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 0:
		goto st_case_0
	}
	goto st_out
tr0:
//line lexer.rl:45
( lex.p) = ( lex.te) - 1
{ tok = int(lex.data[lex.ts]); {( lex.p)++;  lex.cs = 1; goto _out } }
	goto st1
tr2:
//line lexer.rl:44
 lex.te = ( lex.p)+1
{ tok = ENTITY; {( lex.p)++;  lex.cs = 1; goto _out } }
	goto st1
tr3:
//line lexer.rl:45
 lex.te = ( lex.p)+1
{ tok = int(lex.data[lex.ts]); {( lex.p)++;  lex.cs = 1; goto _out } }
	goto st1
tr8:
//line lexer.rl:41
 lex.te = ( lex.p)
( lex.p)--
{ tok = TITLE_VALUE; {( lex.p)++;  lex.cs = 1; goto _out } }
	goto st1
tr10:
//line NONE:1
	switch  lex.act {
	case 1:
	{( lex.p) = ( lex.te) - 1
 tok = TITLE; {( lex.p)++;  lex.cs = 1; goto _out } }
	case 2:
	{( lex.p) = ( lex.te) - 1
 tok = TITLE_VALUE; {( lex.p)++;  lex.cs = 1; goto _out } }
	}
	
	goto st1
tr15:
//line lexer.rl:45
 lex.te = ( lex.p)
( lex.p)--
{ tok = int(lex.data[lex.ts]); {( lex.p)++;  lex.cs = 1; goto _out } }
	goto st1
	st1:
//line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof1
		}
	st_case_1:
//line NONE:1
 lex.ts = ( lex.p)

//line lexer.go:130
		switch  lex.data[( lex.p)] {
		case 32:
			goto st2
		case 84:
			goto st5
		case 91:
			goto tr7
		case 116:
			goto st5
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
				goto st2
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr5
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr5
			}
		default:
			goto tr5
		}
		goto tr3
	st2:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof2
		}
	st_case_2:
		if  lex.data[( lex.p)] == 32 {
			goto st2
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
				goto st2
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto st3
				}
			case  lex.data[( lex.p)] >= 65:
				goto st3
			}
		default:
			goto st3
		}
		goto tr8
	st3:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof3
		}
	st_case_3:
		if  lex.data[( lex.p)] == 32 {
			goto st3
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
				goto st3
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto st3
				}
			case  lex.data[( lex.p)] >= 65:
				goto st3
			}
		default:
			goto st3
		}
		goto tr8
tr5:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:41
 lex.act = 2;
	goto st4
tr14:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:40
 lex.act = 1;
	goto st4
	st4:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof4
		}
	st_case_4:
//line lexer.go:230
		if  lex.data[( lex.p)] == 32 {
			goto st3
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
				goto st3
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr5
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr5
			}
		default:
			goto tr5
		}
		goto tr10
	st5:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof5
		}
	st_case_5:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st3
		case 73:
			goto st6
		case 105:
			goto st6
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
				goto st3
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr5
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr5
			}
		default:
			goto tr5
		}
		goto tr8
	st6:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof6
		}
	st_case_6:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st3
		case 84:
			goto st7
		case 116:
			goto st7
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
				goto st3
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr5
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr5
			}
		default:
			goto tr5
		}
		goto tr8
	st7:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof7
		}
	st_case_7:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st3
		case 76:
			goto st8
		case 108:
			goto st8
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
				goto st3
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr5
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr5
			}
		default:
			goto tr5
		}
		goto tr8
	st8:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof8
		}
	st_case_8:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st3
		case 69:
			goto tr14
		case 101:
			goto tr14
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
				goto st3
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr5
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr5
			}
		default:
			goto tr5
		}
		goto tr8
tr7:
//line NONE:1
 lex.te = ( lex.p)+1

	goto st9
	st9:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof9
		}
	st_case_9:
//line lexer.go:386
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto st0
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st0
			}
		default:
			goto st0
		}
		goto tr15
	st0:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof0
		}
	st_case_0:
		if  lex.data[( lex.p)] == 93 {
			goto tr2
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto st0
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st0
			}
		default:
			goto st0
		}
		goto tr0
	st_out:
	_test_eof1:  lex.cs = 1; goto _test_eof
	_test_eof2:  lex.cs = 2; goto _test_eof
	_test_eof3:  lex.cs = 3; goto _test_eof
	_test_eof4:  lex.cs = 4; goto _test_eof
	_test_eof5:  lex.cs = 5; goto _test_eof
	_test_eof6:  lex.cs = 6; goto _test_eof
	_test_eof7:  lex.cs = 7; goto _test_eof
	_test_eof8:  lex.cs = 8; goto _test_eof
	_test_eof9:  lex.cs = 9; goto _test_eof
	_test_eof0:  lex.cs = 0; goto _test_eof

	_test_eof: {}
	if ( lex.p) == eof {
		switch  lex.cs {
		case 2:
			goto tr8
		case 3:
			goto tr8
		case 4:
			goto tr10
		case 5:
			goto tr8
		case 6:
			goto tr8
		case 7:
			goto tr8
		case 8:
			goto tr8
		case 9:
			goto tr15
		case 0:
			goto tr0
		}
	}

	_out: {}
	}

//line lexer.rl:49


	return tok
}

func (lex *Lexer) string() string {
	return string(lex.data[lex.ts:lex.te])
}
