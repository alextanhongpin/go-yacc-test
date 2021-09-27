
//line lexer.rl:1
package entity_parser


//line lexer.go:7
const lexer_start int = 2
const lexer_first_final int = 2
const lexer_error int = -1

const lexer_en_main int = 2


//line lexer.rl:9



type Attribute struct {
	isPrimary bool
	isForeign bool
	field string
}

type Entity struct {
	name string
	attributes []Attribute
}

type Result []Entity

type Lexer struct {
	data []byte
	p, pe, cs int
	ts, te, act int
	result []Entity
}

func NewLexer(data []byte) *Lexer {
	lex := &Lexer{
		data: data,
		pe: len(data),
	}
	
//line lexer.go:45
	{
	 lex.cs = lexer_start
	 lex.ts = 0
	 lex.te = 0
	 lex.act = 0
	}

//line lexer.rl:38
	return lex
}

func (l *Lexer) Error(msg string) {
	println(msg)
}

func (lex *Lexer) Lex(lval *yySymType) int {
	eof := lex.pe
	var tok int
	yyErrorVerbose = true

	
//line lexer.go:67
	{
	if ( lex.p) == ( lex.pe) {
		goto _test_eof
	}
	switch  lex.cs {
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
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 1:
		goto st_case_1
	}
	goto st_out
tr0:
//line lexer.rl:62
( lex.p) = ( lex.te) - 1
{ lval.str = lex.string(); println("attribute:", lex.string()); tok = ATTRIBUTE; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr1:
//line lexer.rl:61
 lex.te = ( lex.p)+1
{ lval.str = lex.string(); println("string:", lex.string()); tok = STRING; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr2:
//line lexer.rl:68
( lex.p) = ( lex.te) - 1
{ println(string(lex.data[lex.ts])); tok = int(lex.data[lex.ts]); {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr4:
//line lexer.rl:68
 lex.te = ( lex.p)+1
{ println(string(lex.data[lex.ts])); tok = int(lex.data[lex.ts]); {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr8:
//line lexer.rl:59
 lex.te = ( lex.p)+1
{ lval.str = lex.string(); println("*:", lex.string()); tok = PRIMARY_KEY; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr9:
//line lexer.rl:60
 lex.te = ( lex.p)+1
{ lval.str = lex.string(); println("+:", lex.string()); tok = FOREIGN_KEY; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr12:
//line NONE:1
	switch  lex.act {
	case 5:
	{( lex.p) = ( lex.te) - 1
 println("break"); tok = BREAK; {( lex.p)++;  lex.cs = 2; goto _out } }
	case 6:
	{( lex.p) = ( lex.te) - 1
 println("newline"); if lex.te != eof { tok = NEWLINE }; {( lex.p)++;  lex.cs = 2; goto _out } }
	}
	
	goto st2
tr14:
//line lexer.rl:62
 lex.te = ( lex.p)
( lex.p)--
{ lval.str = lex.string(); println("attribute:", lex.string()); tok = ATTRIBUTE; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr20:
//line lexer.rl:61
 lex.te = ( lex.p)
( lex.p)--
{ lval.str = lex.string(); println("string:", lex.string()); tok = STRING; {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
tr21:
//line lexer.rl:68
 lex.te = ( lex.p)
( lex.p)--
{ println(string(lex.data[lex.ts])); tok = int(lex.data[lex.ts]); {( lex.p)++;  lex.cs = 2; goto _out } }
	goto st2
	st2:
//line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof2
		}
	st_case_2:
//line NONE:1
 lex.ts = ( lex.p)

//line lexer.go:170
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr5
		case 13:
			goto tr5
		case 32:
			goto st4
		case 42:
			goto tr8
		case 43:
			goto tr9
		case 96:
			goto tr11
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 45 {
				goto st5
			}
		case  lex.data[( lex.p)] > 90:
			if 95 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st10
			}
		default:
			goto st10
		}
		goto tr4
tr5:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:66
 lex.act = 6;
	goto st3
tr13:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:63
 lex.act = 5;
	goto st3
	st3:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof3
		}
	st_case_3:
//line lexer.go:217
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr13
		case 13:
			goto tr13
		}
		goto tr12
	st4:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof4
		}
	st_case_4:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st5
		case 95:
			goto st5
		case 119:
			goto st6
		}
		switch {
		case  lex.data[( lex.p)] < 44:
			if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 41 {
				goto st5
			}
		case  lex.data[( lex.p)] > 45:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto st5
				}
			case  lex.data[( lex.p)] >= 65:
				goto st5
			}
		default:
			goto st5
		}
		goto tr14
	st5:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof5
		}
	st_case_5:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st5
		case 95:
			goto st5
		}
		switch {
		case  lex.data[( lex.p)] < 44:
			if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 41 {
				goto st5
			}
		case  lex.data[( lex.p)] > 45:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto st5
				}
			case  lex.data[( lex.p)] >= 65:
				goto st5
			}
		default:
			goto st5
		}
		goto tr14
	st6:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof6
		}
	st_case_6:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st5
		case 95:
			goto st5
		case 111:
			goto st7
		}
		switch {
		case  lex.data[( lex.p)] < 44:
			if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 41 {
				goto st5
			}
		case  lex.data[( lex.p)] > 45:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto st5
				}
			case  lex.data[( lex.p)] >= 65:
				goto st5
			}
		default:
			goto st5
		}
		goto tr14
	st7:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof7
		}
	st_case_7:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st5
		case 95:
			goto st5
		case 114:
			goto st8
		}
		switch {
		case  lex.data[( lex.p)] < 44:
			if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 41 {
				goto st5
			}
		case  lex.data[( lex.p)] > 45:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto st5
				}
			case  lex.data[( lex.p)] >= 65:
				goto st5
			}
		default:
			goto st5
		}
		goto tr14
	st8:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof8
		}
	st_case_8:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st5
		case 95:
			goto st5
		case 100:
			goto tr18
		}
		switch {
		case  lex.data[( lex.p)] < 44:
			if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 41 {
				goto st5
			}
		case  lex.data[( lex.p)] > 45:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto st5
				}
			case  lex.data[( lex.p)] >= 65:
				goto st5
			}
		default:
			goto st5
		}
		goto tr14
tr18:
//line NONE:1
 lex.te = ( lex.p)+1

	goto st9
	st9:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof9
		}
	st_case_9:
//line lexer.go:388
		switch  lex.data[( lex.p)] {
		case 32:
			goto st5
		case 43:
			goto st0
		case 95:
			goto st5
		}
		switch {
		case  lex.data[( lex.p)] < 44:
			if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 41 {
				goto st5
			}
		case  lex.data[( lex.p)] > 45:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto st5
				}
			case  lex.data[( lex.p)] >= 65:
				goto st5
			}
		default:
			goto st5
		}
		goto tr14
	st0:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof0
		}
	st_case_0:
		if  lex.data[( lex.p)] == 32 {
			goto tr1
		}
		goto tr0
	st10:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof10
		}
	st_case_10:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st5
		case 95:
			goto st10
		}
		switch {
		case  lex.data[( lex.p)] < 44:
			if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 41 {
				goto st5
			}
		case  lex.data[( lex.p)] > 45:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto st10
				}
			case  lex.data[( lex.p)] >= 65:
				goto st10
			}
		default:
			goto st5
		}
		goto tr20
tr11:
//line NONE:1
 lex.te = ( lex.p)+1

	goto st11
	st11:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof11
		}
	st_case_11:
//line lexer.go:463
		if  lex.data[( lex.p)] == 95 {
			goto st1
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st1
			}
		case  lex.data[( lex.p)] >= 65:
			goto st1
		}
		goto tr21
	st1:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof1
		}
	st_case_1:
		if  lex.data[( lex.p)] == 96 {
			goto tr1
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 95 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st1
			}
		case  lex.data[( lex.p)] >= 65:
			goto st1
		}
		goto tr2
	st_out:
	_test_eof2:  lex.cs = 2; goto _test_eof
	_test_eof3:  lex.cs = 3; goto _test_eof
	_test_eof4:  lex.cs = 4; goto _test_eof
	_test_eof5:  lex.cs = 5; goto _test_eof
	_test_eof6:  lex.cs = 6; goto _test_eof
	_test_eof7:  lex.cs = 7; goto _test_eof
	_test_eof8:  lex.cs = 8; goto _test_eof
	_test_eof9:  lex.cs = 9; goto _test_eof
	_test_eof0:  lex.cs = 0; goto _test_eof
	_test_eof10:  lex.cs = 10; goto _test_eof
	_test_eof11:  lex.cs = 11; goto _test_eof
	_test_eof1:  lex.cs = 1; goto _test_eof

	_test_eof: {}
	if ( lex.p) == eof {
		switch  lex.cs {
		case 3:
			goto tr12
		case 4:
			goto tr14
		case 5:
			goto tr14
		case 6:
			goto tr14
		case 7:
			goto tr14
		case 8:
			goto tr14
		case 9:
			goto tr14
		case 0:
			goto tr0
		case 10:
			goto tr20
		case 11:
			goto tr21
		case 1:
			goto tr2
		}
	}

	_out: {}
	}

//line lexer.rl:73


	return tok
}

func (lex *Lexer) string() string {
	return string(lex.data[lex.ts:lex.te])
}
