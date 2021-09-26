
//line lexer.rl:1
package main


//line lexer.go:7
const lexer_start int = 4
const lexer_first_final int = 4
const lexer_error int = 0

const lexer_en_main int = 4


//line lexer.rl:9


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
	
//line lexer.go:38
	{
	 lex.cs = lexer_start
	 lex.ts = 0
	 lex.te = 0
	 lex.act = 0
	}

//line lexer.rl:31
	return lex
}

func (l *Lexer) Error(msg string) {
	println(msg)
}

func (lex *Lexer) Lex(lval *yySymType) int {
	eof := lex.pe
	var tok int

	
//line lexer.go:59
	{
	if ( lex.p) == ( lex.pe) {
		goto _test_eof
	}
	switch  lex.cs {
	case 4:
		goto st_case_4
	case 0:
		goto st_case_0
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 3:
		goto st_case_3
	}
	goto st_out
tr0:
//line lexer.rl:45
( lex.p) = ( lex.te) - 1
{ lval.school = lex.string(); tok = school; {( lex.p)++;  lex.cs = 4; goto _out } }
	goto st4
tr3:
//line lexer.rl:46
 lex.te = ( lex.p)+1
{ lval.group = lex.string(); tok = group; {( lex.p)++;  lex.cs = 4; goto _out } }
	goto st4
tr5:
//line lexer.rl:48
 lex.te = ( lex.p)+1

	goto st4
tr8:
//line lexer.rl:45
 lex.te = ( lex.p)
( lex.p)--
{ lval.school = lex.string(); tok = school; {( lex.p)++;  lex.cs = 4; goto _out } }
	goto st4
tr11:
//line lexer.rl:44
 lex.te = ( lex.p)
( lex.p)--
{ lval.name = lex.string(); tok = name; {( lex.p)++;  lex.cs = 4; goto _out } }
	goto st4
tr14:
//line lexer.rl:47
 lex.te = ( lex.p)
( lex.p)--
{ lval.module = lex.string(); tok = module; {( lex.p)++;  lex.cs = 4; goto _out } }
	goto st4
tr15:
//line lexer.rl:47
 lex.te = ( lex.p)+1
{ lval.module = lex.string(); tok = module; {( lex.p)++;  lex.cs = 4; goto _out } }
	goto st4
	st4:
//line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof4
		}
	st_case_4:
//line NONE:1
 lex.ts = ( lex.p)

//line lexer.go:136
		switch  lex.data[( lex.p)] {
		case 32:
			goto tr5
		case 95:
			goto st3
		}
		switch {
		case  lex.data[( lex.p)] > 10:
			if 65 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 90 {
				goto st5
			}
		case  lex.data[( lex.p)] >= 9:
			goto tr5
		}
		goto st0
st_case_0:
	st0:
		 lex.cs = 0
		goto _out
	st5:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof5
		}
	st_case_5:
		if  lex.data[( lex.p)] == 39 {
			goto st6
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto tr3
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st6
			}
		default:
			goto tr10
		}
		goto tr8
	st6:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof6
		}
	st_case_6:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st6
		case 39:
			goto st6
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 44 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 45 {
				goto st6
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st6
			}
		default:
			goto st6
		}
		goto tr11
tr10:
//line NONE:1
 lex.te = ( lex.p)+1

	goto st7
	st7:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof7
		}
	st_case_7:
//line lexer.go:211
		switch {
		case  lex.data[( lex.p)] > 57:
			if 65 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 90 {
				goto st9
			}
		case  lex.data[( lex.p)] >= 48:
			goto st1
		}
		goto tr8
	st1:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof1
		}
	st_case_1:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st2
		}
		goto tr0
	st2:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof2
		}
	st_case_2:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st8
		}
		goto tr0
	st8:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof8
		}
	st_case_8:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto tr15
		}
		goto tr14
	st9:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof9
		}
	st_case_9:
		if 65 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 90 {
			goto st9
		}
		goto tr8
	st3:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof3
		}
	st_case_3:
		if  lex.data[( lex.p)] == 95 {
			goto tr3
		}
		goto st0
	st_out:
	_test_eof4:  lex.cs = 4; goto _test_eof
	_test_eof5:  lex.cs = 5; goto _test_eof
	_test_eof6:  lex.cs = 6; goto _test_eof
	_test_eof7:  lex.cs = 7; goto _test_eof
	_test_eof1:  lex.cs = 1; goto _test_eof
	_test_eof2:  lex.cs = 2; goto _test_eof
	_test_eof8:  lex.cs = 8; goto _test_eof
	_test_eof9:  lex.cs = 9; goto _test_eof
	_test_eof3:  lex.cs = 3; goto _test_eof

	_test_eof: {}
	if ( lex.p) == eof {
		switch  lex.cs {
		case 5:
			goto tr8
		case 6:
			goto tr11
		case 7:
			goto tr8
		case 1:
			goto tr0
		case 2:
			goto tr0
		case 8:
			goto tr14
		case 9:
			goto tr8
		}
	}

	_out: {}
	}

//line lexer.rl:52


	return tok
}

func (lex *Lexer) string() string {
	return string(lex.data[lex.ts:lex.te])
}
