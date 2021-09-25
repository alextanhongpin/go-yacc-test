
//line lexer.rl:1
package main


//line lexer.go:7
const lexer_start int = 7
const lexer_first_final int = 7
const lexer_error int = 0

const lexer_en_entity int = 14
const lexer_en_main int = 7


//line lexer.rl:12


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
	
//line lexer.go:35
	{
	 lex.cs = lexer_start
	( lex.top) = 0
	 lex.ts = 0
	 lex.te = 0
	 lex.act = 0
	}

//line lexer.rl:30
	return lex
}

func (l *Lexer) Error(msg string) {
	println(msg)
}

func (lex *Lexer) Lex(lval *yySymType) int {
	eof := lex.pe
	var tok int

	
//line lexer.go:57
	{
	if ( lex.p) == ( lex.pe) {
		goto _test_eof
	}
	switch  lex.cs {
	case 7:
		goto st_case_7
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
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	}
	goto st_out
//line NONE:1
 lex.ts = 0

	st_case_7:
//line lexer.go:109
		switch  lex.data[( lex.p)] {
		case 84:
			goto st1
		case 116:
			goto st1
		}
		goto st0
st_case_0:
	st0:
		 lex.cs = 0
		goto _out
	st1:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof1
		}
	st_case_1:
		switch  lex.data[( lex.p)] {
		case 73:
			goto st2
		case 105:
			goto st2
		}
		goto st0
	st2:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof2
		}
	st_case_2:
		switch  lex.data[( lex.p)] {
		case 84:
			goto st3
		case 116:
			goto st3
		}
		goto st0
	st3:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof3
		}
	st_case_3:
		switch  lex.data[( lex.p)] {
		case 76:
			goto st4
		case 108:
			goto st4
		}
		goto st0
	st4:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof4
		}
	st_case_4:
		switch  lex.data[( lex.p)] {
		case 69:
			goto st5
		case 101:
			goto st5
		}
		goto st0
	st5:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof5
		}
	st_case_5:
		if  lex.data[( lex.p)] == 58 {
			goto st6
		}
		goto st0
	st6:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof6
		}
	st_case_6:
		if  lex.data[( lex.p)] == 32 {
			goto tr6
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
				goto tr6
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr6
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr6
			}
		default:
			goto tr6
		}
		goto st0
tr6:
//line lexer.rl:46

			tok = TITLE;
			{( lex.p)++;  lex.cs = 8; goto _out }
		
//line lexer.rl:57
 {( lex.stack)[( lex.top)] = 8; ( lex.top)++; goto st14 } 
	goto st8
tr8:
//line lexer.rl:57
 {( lex.stack)[( lex.top)] = 8; ( lex.top)++; goto st14 } 
	goto st8
	st8:
//line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof8
		}
	st_case_8:
//line lexer.go:225
		switch  lex.data[( lex.p)] {
		case 32:
			goto tr8
		case 84:
			goto tr9
		case 116:
			goto tr9
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
				goto tr8
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr8
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr8
			}
		default:
			goto tr8
		}
		goto st0
tr9:
//line lexer.rl:57
 {( lex.stack)[( lex.top)] = 9; ( lex.top)++; goto st14 } 
	goto st9
	st9:
//line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof9
		}
	st_case_9:
//line lexer.go:264
		switch  lex.data[( lex.p)] {
		case 32:
			goto tr8
		case 73:
			goto tr10
		case 84:
			goto tr9
		case 105:
			goto tr10
		case 116:
			goto tr9
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
				goto tr8
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr8
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr8
			}
		default:
			goto tr8
		}
		goto st0
tr10:
//line lexer.rl:57
 {( lex.stack)[( lex.top)] = 10; ( lex.top)++; goto st14 } 
	goto st10
	st10:
//line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof10
		}
	st_case_10:
//line lexer.go:307
		switch  lex.data[( lex.p)] {
		case 32:
			goto tr8
		case 84:
			goto tr11
		case 116:
			goto tr11
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
				goto tr8
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr8
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr8
			}
		default:
			goto tr8
		}
		goto st0
tr11:
//line lexer.rl:57
 {( lex.stack)[( lex.top)] = 11; ( lex.top)++; goto st14 } 
	goto st11
	st11:
//line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof11
		}
	st_case_11:
//line lexer.go:346
		switch  lex.data[( lex.p)] {
		case 32:
			goto tr8
		case 73:
			goto tr10
		case 76:
			goto tr12
		case 84:
			goto tr9
		case 105:
			goto tr10
		case 108:
			goto tr12
		case 116:
			goto tr9
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
				goto tr8
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr8
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr8
			}
		default:
			goto tr8
		}
		goto st0
tr12:
//line lexer.rl:57
 {( lex.stack)[( lex.top)] = 12; ( lex.top)++; goto st14 } 
	goto st12
	st12:
//line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof12
		}
	st_case_12:
//line lexer.go:393
		switch  lex.data[( lex.p)] {
		case 32:
			goto tr8
		case 69:
			goto tr13
		case 84:
			goto tr9
		case 101:
			goto tr13
		case 116:
			goto tr9
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
				goto tr8
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr8
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr8
			}
		default:
			goto tr8
		}
		goto st0
tr13:
//line lexer.rl:57
 {( lex.stack)[( lex.top)] = 13; ( lex.top)++; goto st14 } 
	goto st13
	st13:
//line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof13
		}
	st_case_13:
//line lexer.go:436
		switch  lex.data[( lex.p)] {
		case 32:
			goto tr8
		case 58:
			goto st6
		case 84:
			goto tr9
		case 116:
			goto tr9
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
				goto tr8
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr8
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr8
			}
		default:
			goto tr8
		}
		goto st0
tr14:
//line lexer.rl:54
 lex.te = ( lex.p)+1
{ tok = int(lex.data[lex.ts]); {( lex.p)++;  lex.cs = 14; goto _out } }
	goto st14
tr17:
//line NONE:1
	switch  lex.act {
	case 1:
	{( lex.p) = ( lex.te) - 1
 tok = TITLE; {( lex.p)++;  lex.cs = 14; goto _out } }
	case 2:
	{( lex.p) = ( lex.te) - 1
 lval.title = lex.string(); tok = TITLE_VALUE; {( lex.p)++;  lex.cs = 14; goto _out } }
	}
	
	goto st14
tr18:
//line lexer.rl:53
 lex.te = ( lex.p)
( lex.p)--
{ lval.title = lex.string(); tok = TITLE_VALUE; {( lex.p)++;  lex.cs = 14; goto _out } }
	goto st14
	st14:
//line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof14
		}
	st_case_14:
//line NONE:1
 lex.ts = ( lex.p)

//line lexer.go:499
		switch  lex.data[( lex.p)] {
		case 32:
			goto tr15
		case 84:
			goto st16
		case 116:
			goto st16
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
				goto tr15
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr15
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr15
			}
		default:
			goto tr15
		}
		goto tr14
tr15:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:53
 lex.act = 2;
	goto st15
tr22:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:52
 lex.act = 1;
	goto st15
	st15:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof15
		}
	st_case_15:
//line lexer.go:545
		if  lex.data[( lex.p)] == 32 {
			goto tr15
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
				goto tr15
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr15
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr15
			}
		default:
			goto tr15
		}
		goto tr17
	st16:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof16
		}
	st_case_16:
		switch  lex.data[( lex.p)] {
		case 32:
			goto tr15
		case 73:
			goto st17
		case 105:
			goto st17
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
				goto tr15
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr15
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr15
			}
		default:
			goto tr15
		}
		goto tr18
	st17:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof17
		}
	st_case_17:
		switch  lex.data[( lex.p)] {
		case 32:
			goto tr15
		case 84:
			goto st18
		case 116:
			goto st18
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
				goto tr15
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr15
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr15
			}
		default:
			goto tr15
		}
		goto tr18
	st18:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch  lex.data[( lex.p)] {
		case 32:
			goto tr15
		case 76:
			goto st19
		case 108:
			goto st19
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
				goto tr15
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr15
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr15
			}
		default:
			goto tr15
		}
		goto tr18
	st19:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch  lex.data[( lex.p)] {
		case 32:
			goto tr15
		case 69:
			goto tr22
		case 101:
			goto tr22
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
				goto tr15
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto tr15
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr15
			}
		default:
			goto tr15
		}
		goto tr18
	st_out:
	_test_eof1:  lex.cs = 1; goto _test_eof
	_test_eof2:  lex.cs = 2; goto _test_eof
	_test_eof3:  lex.cs = 3; goto _test_eof
	_test_eof4:  lex.cs = 4; goto _test_eof
	_test_eof5:  lex.cs = 5; goto _test_eof
	_test_eof6:  lex.cs = 6; goto _test_eof
	_test_eof8:  lex.cs = 8; goto _test_eof
	_test_eof9:  lex.cs = 9; goto _test_eof
	_test_eof10:  lex.cs = 10; goto _test_eof
	_test_eof11:  lex.cs = 11; goto _test_eof
	_test_eof12:  lex.cs = 12; goto _test_eof
	_test_eof13:  lex.cs = 13; goto _test_eof
	_test_eof14:  lex.cs = 14; goto _test_eof
	_test_eof15:  lex.cs = 15; goto _test_eof
	_test_eof16:  lex.cs = 16; goto _test_eof
	_test_eof17:  lex.cs = 17; goto _test_eof
	_test_eof18:  lex.cs = 18; goto _test_eof
	_test_eof19:  lex.cs = 19; goto _test_eof

	_test_eof: {}
	if ( lex.p) == eof {
		switch  lex.cs {
		case 15:
			goto tr17
		case 16:
			goto tr18
		case 17:
			goto tr18
		case 18:
			goto tr18
		case 19:
			goto tr18
		}
	}

	_out: {}
	}

//line lexer.rl:59


	return tok
}

func (lex *Lexer) string() string {
	return string(lex.data[lex.ts:lex.te])
}
