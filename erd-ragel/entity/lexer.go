
//line lexer.rl:1
package entity_parser


//line lexer.go:7
const lexer_start int = 1
const lexer_first_final int = 23
const lexer_error int = 0

const lexer_en_entity int = 24
const lexer_en_relation int = 27
const lexer_en_main int = 1


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

type Relation struct {
	from string
	to string
	fromCardinality string
	toCardinality string
}

type Result struct {
	entities []Entity
	relations []Relation
}

type Lexer struct {
	data []byte
	p, pe, cs int
	ts, te, act int
	stack []int
	top int

	result Result
}

func NewLexer(data []byte) *Lexer {
	lex := &Lexer{
		data: data,
		pe: len(data),
	}
	
//line lexer.go:60
	{
	 lex.cs = lexer_start
	 lex.top = 0
	 lex.ts = 0
	 lex.te = 0
	 lex.act = 0
	}

//line lexer.rl:51
	return lex
}

func (l *Lexer) Error(msg string) {
	println(msg)
}

func (lex *Lexer) Lex(lval *yySymType) int {
	eof := lex.pe
	var tok int
	yyErrorVerbose = true

	
//line lexer.go:83
	{
	if ( lex.p) == ( lex.pe) {
		goto _test_eof
	}
	goto _resume

_again:
	switch  lex.cs {
	case 1:
		goto st1
	case 23:
		goto st23
	case 24:
		goto st24
	case 0:
		goto st0
	case 25:
		goto st25
	case 26:
		goto st26
	case 2:
		goto st2
	case 3:
		goto st3
	case 4:
		goto st4
	case 5:
		goto st5
	case 6:
		goto st6
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
	case 13:
		goto st13
	case 27:
		goto st27
	case 28:
		goto st28
	case 14:
		goto st14
	case 15:
		goto st15
	case 16:
		goto st16
	case 17:
		goto st17
	case 18:
		goto st18
	case 19:
		goto st19
	case 29:
		goto st29
	case 20:
		goto st20
	case 21:
		goto st21
	case 22:
		goto st22
	}

	if ( lex.p)++; ( lex.p) == ( lex.pe) {
		goto _test_eof
	}
_resume:
	switch  lex.cs {
	case 1:
		goto st_case_1
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 0:
		goto st_case_0
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
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
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
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
	case 29:
		goto st_case_29
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	}
	goto st_out
	st1:
//line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof1
		}
	st_case_1:
//line lexer.go:229
		if  lex.data[( lex.p)] == 91 {
			goto tr1
		}
		goto tr0
tr0:
//line lexer.rl:100

			// fhold does not consume the token.
			( lex.p)--
 {
			// Dynamically resize the stack.
			for len(lex.stack) < lex.top + 1 {
				lex.stack = append(lex.stack, 0)
			}
		{ lex.stack[ lex.top] = 23;  lex.top++; goto st27 }}
		
	goto st23
tr1:
//line lexer.rl:95

			// fhold does not consume the token.
			( lex.p)--
 {
			// Dynamically resize the stack.
			for len(lex.stack) < lex.top + 1 {
				lex.stack = append(lex.stack, 0)
			}
		{ lex.stack[ lex.top] = 23;  lex.top++; goto st24 }}
		
	goto st23
	st23:
//line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof23
		}
	st_case_23:
//line lexer.go:268
		if  lex.data[( lex.p)] == 91 {
			goto tr1
		}
		goto tr0
tr10:
//line lexer.rl:79
 lex.te = ( lex.p)+1
{ tok = ENTITY; lval.str = string(lex.data[lex.ts+1:lex.te-1]); {( lex.p)++;  lex.cs = 24; goto _out } }
	goto st24
tr23:
//line lexer.rl:80
 lex.te = ( lex.p)+1
{ tok = PRIMARY_KEY; lval.str = lex.string(); {( lex.p)++;  lex.cs = 24; goto _out } }
	goto st24
tr24:
//line lexer.rl:81
 lex.te = ( lex.p)+1
{ tok = FOREIGN_KEY; lval.str = lex.string(); {( lex.p)++;  lex.cs = 24; goto _out } }
	goto st24
tr27:
//line NONE:1
	switch  lex.act {
	case 5:
	{( lex.p) = ( lex.te) - 1
 { lex.top--;  lex.cs =  lex.stack[ lex.top];goto _again } }
	default:
	{( lex.p) = ( lex.te) - 1
}
	}
	
	goto st24
tr29:
//line lexer.rl:82
 lex.te = ( lex.p)
( lex.p)--
{ tok = ATTRIBUTE; lval.str = lex.string(); {( lex.p)++;  lex.cs = 24; goto _out } }
	goto st24
	st24:
//line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof24
		}
	st_case_24:
//line NONE:1
 lex.ts = ( lex.p)

//line lexer.go:317
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr22
		case 13:
			goto tr22
		case 42:
			goto tr23
		case 43:
			goto tr24
		case 91:
			goto st2
		case 95:
			goto st26
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st26
			}
		case  lex.data[( lex.p)] >= 65:
			goto st26
		}
		goto st0
st_case_0:
	st0:
		 lex.cs = 0
		goto _out
tr22:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:84
 lex.act = 6;
	goto st25
tr28:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:83
 lex.act = 5;
	goto st25
	st25:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof25
		}
	st_case_25:
//line lexer.go:364
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr28
		case 13:
			goto tr28
		}
		goto tr27
	st26:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof26
		}
	st_case_26:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st26
		case 95:
			goto st26
		}
		switch {
		case  lex.data[( lex.p)] < 44:
			if 40 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 41 {
				goto st26
			}
		case  lex.data[( lex.p)] > 45:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto st26
				}
			case  lex.data[( lex.p)] >= 65:
				goto st26
			}
		default:
			goto st26
		}
		goto tr29
	st2:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof2
		}
	st_case_2:
		switch  lex.data[( lex.p)] {
		case 34:
			goto st3
		case 39:
			goto st7
		case 96:
			goto st11
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 95 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st10
			}
		case  lex.data[( lex.p)] >= 65:
			goto st10
		}
		goto st0
	st3:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof3
		}
	st_case_3:
		if  lex.data[( lex.p)] == 95 {
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
		goto st0
	st4:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof4
		}
	st_case_4:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st4
		case 95:
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
		goto st0
	st5:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof5
		}
	st_case_5:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st4
		case 34:
			goto st6
		case 95:
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
		goto st0
	st6:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof6
		}
	st_case_6:
		if  lex.data[( lex.p)] == 93 {
			goto tr10
		}
		goto st0
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
		case 39:
			goto st6
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
		switch  lex.data[( lex.p)] {
		case 93:
			goto tr10
		case 95:
			goto st10
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st10
			}
		case  lex.data[( lex.p)] >= 65:
			goto st10
		}
		goto st0
	st11:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof11
		}
	st_case_11:
		if  lex.data[( lex.p)] == 95 {
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
			goto st12
		case 95:
			goto st13
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st13
			}
		case  lex.data[( lex.p)] >= 65:
			goto st13
		}
		goto st0
	st13:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof13
		}
	st_case_13:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st12
		case 96:
			goto st6
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 95 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st13
			}
		case  lex.data[( lex.p)] >= 65:
			goto st13
		}
		goto st0
tr17:
//line lexer.rl:89
 lex.te = ( lex.p)+1
{ tok = ENTITY; lval.str = lex.string(); {( lex.p)++;  lex.cs = 27; goto _out } }
	goto st27
tr31:
//line lexer.rl:91
 lex.te = ( lex.p)+1

	goto st27
tr34:
//line lexer.rl:88
 lex.te = ( lex.p)+1
{ tok = CARDINALITY; lval.str = lex.string(); {( lex.p)++;  lex.cs = 27; goto _out } }
	goto st27
tr37:
//line NONE:1
	switch  lex.act {
	case 9:
	{( lex.p) = ( lex.te) - 1
 { lex.top--;  lex.cs =  lex.stack[ lex.top];goto _again } }
	default:
	{( lex.p) = ( lex.te) - 1
}
	}
	
	goto st27
tr39:
//line lexer.rl:89
 lex.te = ( lex.p)
( lex.p)--
{ tok = ENTITY; lval.str = lex.string(); {( lex.p)++;  lex.cs = 27; goto _out } }
	goto st27
	st27:
//line NONE:1
 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof27
		}
	st_case_27:
//line NONE:1
 lex.ts = ( lex.p)

//line lexer.go:671
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr30
		case 13:
			goto tr30
		case 32:
			goto tr31
		case 34:
			goto st14
		case 39:
			goto st17
		case 45:
			goto tr31
		case 49:
			goto tr34
		case 63:
			goto tr34
		case 96:
			goto st20
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 42 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 43 {
				goto tr34
			}
		case  lex.data[( lex.p)] > 90:
			if 95 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st29
			}
		default:
			goto st29
		}
		goto st0
tr30:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:92
 lex.act = 11;
	goto st28
tr38:
//line NONE:1
 lex.te = ( lex.p)+1

//line lexer.rl:90
 lex.act = 9;
	goto st28
	st28:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof28
		}
	st_case_28:
//line lexer.go:724
		switch  lex.data[( lex.p)] {
		case 10:
			goto tr38
		case 13:
			goto tr38
		}
		goto tr37
	st14:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof14
		}
	st_case_14:
		if  lex.data[( lex.p)] == 95 {
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
			goto st15
		case 95:
			goto st16
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st16
			}
		case  lex.data[( lex.p)] >= 65:
			goto st16
		}
		goto st0
	st16:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof16
		}
	st_case_16:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st15
		case 34:
			goto tr17
		case 95:
			goto st16
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st16
			}
		case  lex.data[( lex.p)] >= 65:
			goto st16
		}
		goto st0
	st17:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof17
		}
	st_case_17:
		if  lex.data[( lex.p)] == 95 {
			goto st18
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st18
			}
		case  lex.data[( lex.p)] >= 65:
			goto st18
		}
		goto st0
	st18:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st18
		case 95:
			goto st19
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st19
			}
		case  lex.data[( lex.p)] >= 65:
			goto st19
		}
		goto st0
	st19:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st18
		case 39:
			goto tr17
		case 95:
			goto st19
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st19
			}
		case  lex.data[( lex.p)] >= 65:
			goto st19
		}
		goto st0
	st29:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof29
		}
	st_case_29:
		if  lex.data[( lex.p)] == 95 {
			goto st29
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st29
			}
		case  lex.data[( lex.p)] >= 65:
			goto st29
		}
		goto tr39
	st20:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof20
		}
	st_case_20:
		if  lex.data[( lex.p)] == 95 {
			goto st21
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st21
			}
		case  lex.data[( lex.p)] >= 65:
			goto st21
		}
		goto st0
	st21:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof21
		}
	st_case_21:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st21
		case 95:
			goto st22
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st22
			}
		case  lex.data[( lex.p)] >= 65:
			goto st22
		}
		goto st0
	st22:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof22
		}
	st_case_22:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st21
		case 96:
			goto tr17
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 95 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st22
			}
		case  lex.data[( lex.p)] >= 65:
			goto st22
		}
		goto st0
	st_out:
	_test_eof1:  lex.cs = 1; goto _test_eof
	_test_eof23:  lex.cs = 23; goto _test_eof
	_test_eof24:  lex.cs = 24; goto _test_eof
	_test_eof25:  lex.cs = 25; goto _test_eof
	_test_eof26:  lex.cs = 26; goto _test_eof
	_test_eof2:  lex.cs = 2; goto _test_eof
	_test_eof3:  lex.cs = 3; goto _test_eof
	_test_eof4:  lex.cs = 4; goto _test_eof
	_test_eof5:  lex.cs = 5; goto _test_eof
	_test_eof6:  lex.cs = 6; goto _test_eof
	_test_eof7:  lex.cs = 7; goto _test_eof
	_test_eof8:  lex.cs = 8; goto _test_eof
	_test_eof9:  lex.cs = 9; goto _test_eof
	_test_eof10:  lex.cs = 10; goto _test_eof
	_test_eof11:  lex.cs = 11; goto _test_eof
	_test_eof12:  lex.cs = 12; goto _test_eof
	_test_eof13:  lex.cs = 13; goto _test_eof
	_test_eof27:  lex.cs = 27; goto _test_eof
	_test_eof28:  lex.cs = 28; goto _test_eof
	_test_eof14:  lex.cs = 14; goto _test_eof
	_test_eof15:  lex.cs = 15; goto _test_eof
	_test_eof16:  lex.cs = 16; goto _test_eof
	_test_eof17:  lex.cs = 17; goto _test_eof
	_test_eof18:  lex.cs = 18; goto _test_eof
	_test_eof19:  lex.cs = 19; goto _test_eof
	_test_eof29:  lex.cs = 29; goto _test_eof
	_test_eof20:  lex.cs = 20; goto _test_eof
	_test_eof21:  lex.cs = 21; goto _test_eof
	_test_eof22:  lex.cs = 22; goto _test_eof

	_test_eof: {}
	if ( lex.p) == eof {
		switch  lex.cs {
		case 25:
			goto tr27
		case 26:
			goto tr29
		case 28:
			goto tr37
		case 29:
			goto tr39
		}
	}

	_out: {}
	}

//line lexer.rl:115


	return tok
}

func (lex *Lexer) string() string {
	return string(lex.data[lex.ts:lex.te])
}

func (lex *Lexer) token() int {
	return int(lex.data[lex.ts])
}
