
//line lexer.rl:1
package entity_parser


//line lexer.go:7
const lexer_start int = 16
const lexer_first_final int = 16
const lexer_error int = 0

const lexer_en_entity int = 23
const lexer_en_main int = 16


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
	stack []int
	top int
	result []Entity
}

func NewLexer(data []byte) *Lexer {
	lex := &Lexer{
		data: data,
		pe: len(data),
	}
	
//line lexer.go:48
	{
	 lex.cs = lexer_start
	 lex.top = 0
	 lex.ts = 0
	 lex.te = 0
	 lex.act = 0
	}

//line lexer.rl:40
	return lex
}

func (l *Lexer) Error(msg string) {
	println(msg)
}

func (lex *Lexer) Lex(lval *yySymType) int {
	eof := lex.pe
	var tok int
	yyErrorVerbose = true

	
//line lexer.go:71
	{
	if ( lex.p) == ( lex.pe) {
		goto _test_eof
	}
	goto _resume

_again:
	switch  lex.cs {
	case 16:
		goto st16
	case 17:
		goto st17
	case 18:
		goto st18
	case 19:
		goto st19
	case 1:
		goto st1
	case 2:
		goto st2
	case 20:
		goto st20
	case 3:
		goto st3
	case 4:
		goto st4
	case 21:
		goto st21
	case 22:
		goto st22
	case 5:
		goto st5
	case 6:
		goto st6
	case 23:
		goto st23
	case 0:
		goto st0
	case 7:
		goto st7
	case 8:
		goto st8
	case 9:
		goto st9
	case 10:
		goto st10
	case 11:
		goto st11
	case 12:
		goto st12
	case 24:
		goto st24
	case 13:
		goto st13
	case 14:
		goto st14
	case 15:
		goto st15
	}

	if ( lex.p)++; ( lex.p) == ( lex.pe) {
		goto _test_eof
	}
_resume:
	switch  lex.cs {
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 20:
		goto st_case_20
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 23:
		goto st_case_23
	case 0:
		goto st_case_0
	case 7:
		goto st_case_7
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
	case 24:
		goto st_case_24
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	}
	goto st_out
tr0:
//line lexer.rl:83
( lex.p) = ( lex.te) - 1
{ tok = lex.token(); {( lex.p)++;  lex.cs = 16; goto _out } }
	goto st16
tr3:
//line lexer.rl:75
 lex.te = ( lex.p)+1
{ lval.str = lex.string(); println("string:", lex.string()); tok = STRING; {( lex.p)++;  lex.cs = 16; goto _out } }
	goto st16
tr16:
//line lexer.rl:83
 lex.te = ( lex.p)+1
{ tok = lex.token(); {( lex.p)++;  lex.cs = 16; goto _out } }
	goto st16
tr21:
//line lexer.rl:73
 lex.te = ( lex.p)+1
{ lval.str = lex.string(); println("*:", lex.string()); tok = PRIMARY_KEY; {( lex.p)++;  lex.cs = 16; goto _out } }
	goto st16
tr22:
//line lexer.rl:74
 lex.te = ( lex.p)+1
{ lval.str = lex.string(); println("+:", lex.string()); tok = FOREIGN_KEY; {( lex.p)++;  lex.cs = 16; goto _out } }
	goto st16
tr24:
//line lexer.rl:81
 lex.te = ( lex.p)+1
{ ( lex.p)--
 {
			// Dynamically resize the stack.
			for len(lex.stack) < lex.top + 1 {
				lex.stack = append(lex.stack, 0)
			}
		{ lex.stack[ lex.top] = 16;  lex.top++; goto st23 }} }
	goto st16
tr26:
//line NONE:1
	switch  lex.act {
	case 8:
	{( lex.p) = ( lex.te) - 1
 tok = BREAK; {( lex.p)++;  lex.cs = 16; goto _out } }
	case 9:
	{( lex.p) = ( lex.te) - 1
 if lex.te != eof { tok = NEWLINE }; {( lex.p)++;  lex.cs = 16; goto _out } }
	}
	
	goto st16
tr28:
//line lexer.rl:76
 lex.te = ( lex.p)
( lex.p)--
{ lval.str = lex.string(); tok = ATTRIBUTE; {( lex.p)++;  lex.cs = 16; goto _out } }
	goto st16
tr29:
//line lexer.rl:83
 lex.te = ( lex.p)
( lex.p)--
{ tok = lex.token(); {( lex.p)++;  lex.cs = 16; goto _out } }
	goto st16
tr30:
//line lexer.rl:75
 lex.te = ( lex.p)
( lex.p)--
{ lval.str = lex.string(); println("string:", lex.string()); tok = STRING; {( lex.p)++;  lex.cs = 16; goto _out } }
	goto st16
	st16:
//line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof16
		}
	st_case_16:
//line NONE:1
 lex.ts = ( lex.p)

//line lexer.go:266
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr17
		case 13:
			goto tr17
		case 32:
			goto st18
		case 34:
			goto tr19
		case 39:
			goto tr20
		case 42:
			goto tr21
		case 43:
			goto tr22
		case 91:
			goto tr24
		case 96:
			goto tr25
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 45 {
				goto st18
			}
		case  lex.data[( lex.p)] > 90:
			if 95 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st21
			}
		default:
			goto st21
		}
		goto tr16
tr17:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:80
 lex.act = 9;
	goto st17
tr27:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:77
 lex.act = 8;
	goto st17
	st17:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof17
		}
	st_case_17:
//line lexer.go:319
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr27
		case 13:
			goto tr27
		}
		goto tr26
	st18:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st18
		case 95:
			goto st18
		}
		switch {
		case  lex.data[( lex.p)] < 44:
			if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 41 {
				goto st18
			}
		case  lex.data[( lex.p)] > 45:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto st18
				}
			case  lex.data[( lex.p)] >= 65:
				goto st18
			}
		default:
			goto st18
		}
		goto tr28
tr19:
//line NONE:1
 lex.te = ( lex.p)+1

	goto st19
	st19:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof19
		}
	st_case_19:
//line lexer.go:366
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
		goto tr29
	st1:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof1
		}
	st_case_1:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st1
		case 95:
			goto st2
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st2
			}
		case  lex.data[( lex.p)] >= 65:
			goto st2
		}
		goto tr0
	st2:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof2
		}
	st_case_2:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st1
		case 34:
			goto tr3
		case 95:
			goto st2
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st2
			}
		case  lex.data[( lex.p)] >= 65:
			goto st2
		}
		goto tr0
tr20:
//line NONE:1
 lex.te = ( lex.p)+1

	goto st20
	st20:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof20
		}
	st_case_20:
//line lexer.go:431
		if  lex.data[( lex.p)] == 95 {
			goto st3
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st3
			}
		case  lex.data[( lex.p)] >= 65:
			goto st3
		}
		goto tr29
	st3:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof3
		}
	st_case_3:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st3
		case 95:
			goto st4
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st4
			}
		case  lex.data[( lex.p)] >= 65:
			goto st4
		}
		goto tr0
	st4:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof4
		}
	st_case_4:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st3
		case 39:
			goto tr3
		case 95:
			goto st4
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st4
			}
		case  lex.data[( lex.p)] >= 65:
			goto st4
		}
		goto tr0
	st21:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof21
		}
	st_case_21:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st18
		case 95:
			goto st21
		}
		switch {
		case  lex.data[( lex.p)] < 44:
			if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 41 {
				goto st18
			}
		case  lex.data[( lex.p)] > 45:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto st21
				}
			case  lex.data[( lex.p)] >= 65:
				goto st21
			}
		default:
			goto st18
		}
		goto tr30
tr25:
//line NONE:1
 lex.te = ( lex.p)+1

	goto st22
	st22:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof22
		}
	st_case_22:
//line lexer.go:525
		if  lex.data[( lex.p)] == 95 {
			goto st5
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st5
			}
		case  lex.data[( lex.p)] >= 65:
			goto st5
		}
		goto tr29
	st5:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof5
		}
	st_case_5:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st5
		case 95:
			goto st6
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st6
			}
		case  lex.data[( lex.p)] >= 65:
			goto st6
		}
		goto tr0
	st6:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof6
		}
	st_case_6:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st5
		case 96:
			goto tr3
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 95 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st6
			}
		case  lex.data[( lex.p)] >= 65:
			goto st6
		}
		goto tr0
tr11:
//line lexer.rl:67
 lex.te = ( lex.p)+1
{ tok = STRING; lval.str = lex.string(); {( lex.p)++;  lex.cs = 23; goto _out } }
	goto st23
tr31:
//line lexer.rl:69
 lex.te = ( lex.p)+1
{ ( lex.p)--
 { lex.top--;  lex.cs =  lex.stack[ lex.top];goto _again } }
	goto st23
tr35:
//line lexer.rl:68
 lex.te = ( lex.p)+1
{ tok = lex.token(); {( lex.p)++;  lex.cs = 23; goto _out } }
	goto st23
tr37:
//line lexer.rl:67
 lex.te = ( lex.p)
( lex.p)--
{ tok = STRING; lval.str = lex.string(); {( lex.p)++;  lex.cs = 23; goto _out } }
	goto st23
	st23:
//line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof23
		}
	st_case_23:
//line NONE:1
 lex.ts = ( lex.p)

//line lexer.go:611
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr31
		case 13:
			goto tr31
		case 34:
			goto st7
		case 39:
			goto st10
		case 91:
			goto tr35
		case 93:
			goto tr35
		case 96:
			goto st13
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 95 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st24
			}
		case  lex.data[( lex.p)] >= 65:
			goto st24
		}
		goto st0
st_case_0:
	st0:
		 lex.cs = 0
		goto _out
	st7:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof7
		}
	st_case_7:
		if  lex.data[( lex.p)] == 95 {
			goto st8
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st8
			}
		case  lex.data[( lex.p)] >= 65:
			goto st8
		}
		goto st0
	st8:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof8
		}
	st_case_8:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st8
		case 95:
			goto st9
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st9
			}
		case  lex.data[( lex.p)] >= 65:
			goto st9
		}
		goto st0
	st9:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof9
		}
	st_case_9:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st8
		case 34:
			goto tr11
		case 95:
			goto st9
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st9
			}
		case  lex.data[( lex.p)] >= 65:
			goto st9
		}
		goto st0
	st10:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof10
		}
	st_case_10:
		if  lex.data[( lex.p)] == 95 {
			goto st11
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st11
			}
		case  lex.data[( lex.p)] >= 65:
			goto st11
		}
		goto st0
	st11:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof11
		}
	st_case_11:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st11
		case 95:
			goto st12
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st12
			}
		case  lex.data[( lex.p)] >= 65:
			goto st12
		}
		goto st0
	st12:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof12
		}
	st_case_12:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st11
		case 39:
			goto tr11
		case 95:
			goto st12
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st12
			}
		case  lex.data[( lex.p)] >= 65:
			goto st12
		}
		goto st0
	st24:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof24
		}
	st_case_24:
		if  lex.data[( lex.p)] == 95 {
			goto st24
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st24
			}
		case  lex.data[( lex.p)] >= 65:
			goto st24
		}
		goto tr37
	st13:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof13
		}
	st_case_13:
		if  lex.data[( lex.p)] == 95 {
			goto st14
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st14
			}
		case  lex.data[( lex.p)] >= 65:
			goto st14
		}
		goto st0
	st14:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof14
		}
	st_case_14:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st14
		case 95:
			goto st15
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st15
			}
		case  lex.data[( lex.p)] >= 65:
			goto st15
		}
		goto st0
	st15:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof15
		}
	st_case_15:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st14
		case 96:
			goto tr11
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 95 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st15
			}
		case  lex.data[( lex.p)] >= 65:
			goto st15
		}
		goto st0
	st_out:
	_test_eof16:  lex.cs = 16; goto _test_eof
	_test_eof17:  lex.cs = 17; goto _test_eof
	_test_eof18:  lex.cs = 18; goto _test_eof
	_test_eof19:  lex.cs = 19; goto _test_eof
	_test_eof1:  lex.cs = 1; goto _test_eof
	_test_eof2:  lex.cs = 2; goto _test_eof
	_test_eof20:  lex.cs = 20; goto _test_eof
	_test_eof3:  lex.cs = 3; goto _test_eof
	_test_eof4:  lex.cs = 4; goto _test_eof
	_test_eof21:  lex.cs = 21; goto _test_eof
	_test_eof22:  lex.cs = 22; goto _test_eof
	_test_eof5:  lex.cs = 5; goto _test_eof
	_test_eof6:  lex.cs = 6; goto _test_eof
	_test_eof23:  lex.cs = 23; goto _test_eof
	_test_eof7:  lex.cs = 7; goto _test_eof
	_test_eof8:  lex.cs = 8; goto _test_eof
	_test_eof9:  lex.cs = 9; goto _test_eof
	_test_eof10:  lex.cs = 10; goto _test_eof
	_test_eof11:  lex.cs = 11; goto _test_eof
	_test_eof12:  lex.cs = 12; goto _test_eof
	_test_eof24:  lex.cs = 24; goto _test_eof
	_test_eof13:  lex.cs = 13; goto _test_eof
	_test_eof14:  lex.cs = 14; goto _test_eof
	_test_eof15:  lex.cs = 15; goto _test_eof

	_test_eof: {}
	if ( lex.p) == eof {
		switch  lex.cs {
		case 17:
			goto tr26
		case 18:
			goto tr28
		case 19:
			goto tr29
		case 1:
			goto tr0
		case 2:
			goto tr0
		case 20:
			goto tr29
		case 3:
			goto tr0
		case 4:
			goto tr0
		case 21:
			goto tr30
		case 22:
			goto tr29
		case 5:
			goto tr0
		case 6:
			goto tr0
		case 24:
			goto tr37
		}
	}

	_out: {}
	}

//line lexer.rl:88


	return tok
}

func (lex *Lexer) string() string {
	return string(lex.data[lex.ts:lex.te])
}

func (lex *Lexer) token() int {
	return int(lex.data[lex.ts])
}
