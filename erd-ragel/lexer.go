
//line lexer.rl:1
package main


//line lexer.go:7
const lexer_start int = 3
const lexer_first_final int = 3
const lexer_error int = -1

const lexer_en_main int = 3


//line lexer.rl:9


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
	
//line lexer.go:50
	{
	 lex.cs = lexer_start
	 lex.ts = 0
	 lex.te = 0
	 lex.act = 0
	}

//line lexer.rl:43
	return lex
}

func (l *Lexer) Error(msg string) {
	println(msg)
}

func (lex *Lexer) Lex(lval *yySymType) int {
	eof := lex.pe
	var tok int
	yyErrorVerbose = true

	
//line lexer.go:72
	{
	if ( lex.p) == ( lex.pe) {
		goto _test_eof
	}
	switch  lex.cs {
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 0:
		goto st_case_0
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
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	}
	goto st_out
tr0:
//line NONE:1
	switch  lex.act {
	case 2:
	{( lex.p) = ( lex.te) - 1
 tok = TITLE; {( lex.p)++;  lex.cs = 3; goto _out } }
	case 3:
	{( lex.p) = ( lex.te) - 1
 lval.str = lex.string(); println("string:", lex.string()); tok = STRING; {( lex.p)++;  lex.cs = 3; goto _out } }
	case 4:
	{( lex.p) = ( lex.te) - 1
 println("break"); tok = BREAK; {( lex.p)++;  lex.cs = 3; goto _out } }
	case 5:
	{( lex.p) = ( lex.te) - 1
 println("newline", lex.te, eof); if lex.te != eof { tok = NEWLINE }; {( lex.p)++;  lex.cs = 3; goto _out } }
	case 7:
	{( lex.p) = ( lex.te) - 1
 println(string(lex.data[lex.ts])); tok = int(lex.data[lex.ts]); {( lex.p)++;  lex.cs = 3; goto _out } }
	}
	
	goto st3
tr3:
//line lexer.rl:74
( lex.p) = ( lex.te) - 1
{ println(string(lex.data[lex.ts])); tok = int(lex.data[lex.ts]); {( lex.p)++;  lex.cs = 3; goto _out } }
	goto st3
tr5:
//line lexer.rl:68
 lex.te = ( lex.p)+1
{ lval.str = lex.string(); println("string:", lex.string()); tok = STRING; {( lex.p)++;  lex.cs = 3; goto _out } }
	goto st3
tr6:
//line lexer.rl:74
 lex.te = ( lex.p)+1
{ println(string(lex.data[lex.ts])); tok = int(lex.data[lex.ts]); {( lex.p)++;  lex.cs = 3; goto _out } }
	goto st3
tr15:
//line lexer.rl:73
 lex.te = ( lex.p)
( lex.p)--

	goto st3
tr16:
//line lexer.rl:66
 lex.te = ( lex.p)+1
{ println("cardinality", lex.string()); lval.str = lex.string(); tok = CARDINALITY; {( lex.p)++;  lex.cs = 3; goto _out } }
	goto st3
tr17:
//line lexer.rl:74
 lex.te = ( lex.p)
( lex.p)--
{ println(string(lex.data[lex.ts])); tok = int(lex.data[lex.ts]); {( lex.p)++;  lex.cs = 3; goto _out } }
	goto st3
tr20:
//line lexer.rl:68
 lex.te = ( lex.p)
( lex.p)--
{ lval.str = lex.string(); println("string:", lex.string()); tok = STRING; {( lex.p)++;  lex.cs = 3; goto _out } }
	goto st3
	st3:
//line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof3
		}
	st_case_3:
//line NONE:1
 lex.ts = ( lex.p)

//line lexer.go:178
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr7
		case 13:
			goto tr7
		case 32:
			goto st5
		case 49:
			goto st8
		case 63:
			goto st8
		case 84:
			goto tr12
		case 96:
			goto tr13
		case 116:
			goto tr12
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 42 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 43 {
				goto tr9
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr6
tr7:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:72
 lex.act = 5;
	goto st4
tr14:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:69
 lex.act = 4;
	goto st4
	st4:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof4
		}
	st_case_4:
//line lexer.go:229
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr14
		case 13:
			goto tr14
		}
		goto tr0
	st5:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof5
		}
	st_case_5:
		switch  lex.data[( lex.p)] {
		case 49:
			goto tr16
		case 63:
			goto tr16
		}
		if 42 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 43 {
			goto tr16
		}
		goto tr15
tr9:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:74
 lex.act = 7;
	goto st6
	st6:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof6
		}
	st_case_6:
//line lexer.go:264
		switch  lex.data[( lex.p)] {
		case 32:
			goto tr16
		case 96:
			goto st1
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st0
			}
		case  lex.data[( lex.p)] >= 65:
			goto st0
		}
		goto tr17
	st0:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof0
		}
	st_case_0:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st0
		case 44:
			goto tr2
		case 95:
			goto tr2
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 41 {
				goto tr2
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr2
			}
		default:
			goto tr2
		}
		goto tr0
tr2:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:68
 lex.act = 3;
	goto st7
tr11:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:74
 lex.act = 7;
	goto st7
tr23:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:67
 lex.act = 2;
	goto st7
	st7:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof7
		}
	st_case_7:
//line lexer.go:332
		switch  lex.data[( lex.p)] {
		case 32:
			goto st0
		case 44:
			goto tr2
		case 95:
			goto tr2
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 41 {
				goto tr2
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr2
			}
		default:
			goto tr2
		}
		goto tr0
	st1:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof1
		}
	st_case_1:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st2
		case 44:
			goto st2
		case 95:
			goto st2
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 41 {
				goto st2
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st2
			}
		default:
			goto st2
		}
		goto tr3
	st2:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof2
		}
	st_case_2:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st2
		case 44:
			goto st2
		case 96:
			goto tr5
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 41 {
				goto st2
			}
		case  lex.data[( lex.p)] > 90:
			if 95 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st2
			}
		default:
			goto st2
		}
		goto tr3
	st8:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof8
		}
	st_case_8:
		if  lex.data[( lex.p)] == 32 {
			goto tr16
		}
		goto tr17
tr12:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:74
 lex.act = 7;
	goto st9
	st9:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof9
		}
	st_case_9:
//line lexer.go:427
		switch  lex.data[( lex.p)] {
		case 32:
			goto st0
		case 44:
			goto tr2
		case 73:
			goto tr19
		case 95:
			goto tr2
		case 105:
			goto tr19
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 41 {
				goto tr2
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr2
			}
		default:
			goto tr2
		}
		goto tr17
tr19:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:68
 lex.act = 3;
	goto st10
	st10:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof10
		}
	st_case_10:
//line lexer.go:465
		switch  lex.data[( lex.p)] {
		case 32:
			goto st0
		case 44:
			goto tr2
		case 84:
			goto tr21
		case 95:
			goto tr2
		case 116:
			goto tr21
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 41 {
				goto tr2
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr2
			}
		default:
			goto tr2
		}
		goto tr20
tr21:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:68
 lex.act = 3;
	goto st11
	st11:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof11
		}
	st_case_11:
//line lexer.go:503
		switch  lex.data[( lex.p)] {
		case 32:
			goto st0
		case 44:
			goto tr2
		case 76:
			goto tr22
		case 95:
			goto tr2
		case 108:
			goto tr22
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 41 {
				goto tr2
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr2
			}
		default:
			goto tr2
		}
		goto tr20
tr22:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:68
 lex.act = 3;
	goto st12
	st12:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof12
		}
	st_case_12:
//line lexer.go:541
		switch  lex.data[( lex.p)] {
		case 32:
			goto st0
		case 44:
			goto tr2
		case 69:
			goto tr23
		case 95:
			goto tr2
		case 101:
			goto tr23
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 41 {
				goto tr2
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto tr2
			}
		default:
			goto tr2
		}
		goto tr20
tr13:
//line NONE:1
 lex.te = ( lex.p)+1

	goto st13
	st13:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof13
		}
	st_case_13:
//line lexer.go:577
		switch  lex.data[( lex.p)] {
		case 32:
			goto st2
		case 44:
			goto st2
		case 95:
			goto st2
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 41 {
				goto st2
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st2
			}
		default:
			goto st2
		}
		goto tr17
	st_out:
	_test_eof3:  lex.cs = 3; goto _test_eof
	_test_eof4:  lex.cs = 4; goto _test_eof
	_test_eof5:  lex.cs = 5; goto _test_eof
	_test_eof6:  lex.cs = 6; goto _test_eof
	_test_eof0:  lex.cs = 0; goto _test_eof
	_test_eof7:  lex.cs = 7; goto _test_eof
	_test_eof1:  lex.cs = 1; goto _test_eof
	_test_eof2:  lex.cs = 2; goto _test_eof
	_test_eof8:  lex.cs = 8; goto _test_eof
	_test_eof9:  lex.cs = 9; goto _test_eof
	_test_eof10:  lex.cs = 10; goto _test_eof
	_test_eof11:  lex.cs = 11; goto _test_eof
	_test_eof12:  lex.cs = 12; goto _test_eof
	_test_eof13:  lex.cs = 13; goto _test_eof

	_test_eof: {}
	if ( lex.p) == eof {
		switch  lex.cs {
		case 4:
			goto tr0
		case 5:
			goto tr15
		case 6:
			goto tr17
		case 0:
			goto tr0
		case 7:
			goto tr0
		case 1:
			goto tr3
		case 2:
			goto tr3
		case 8:
			goto tr17
		case 9:
			goto tr17
		case 10:
			goto tr20
		case 11:
			goto tr20
		case 12:
			goto tr20
		case 13:
			goto tr17
		}
	}

	_out: {}
	}

//line lexer.rl:79


	return tok
}

func (lex *Lexer) string() string {
	return string(lex.data[lex.ts:lex.te])
}
