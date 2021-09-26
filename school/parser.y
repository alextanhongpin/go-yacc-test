%{
package main

func setResult(l yyLexer, val Student) {
	l.(*Lexer).result = []Student{val}
}

func appendResult(l yyLexer, val Student) {
	l.(*Lexer).result = append(l.(*Lexer).result, val)
}
%}

%union {
	school string
	group string
	module string
	modules []string
	name string
	student Student
}

%token <school> school
%token <group> group
%token <module> module
%token <name> name

%type <modules> modules
%type <student> student file

%start file

%%

file: student { setResult(yylex, $$); }
		| file student { appendResult(yylex, $2); }
		;

student	: school group name modules
				{
					$$.school = $1;
					$$.group = $2;
					$$.name = $3;
					$$.modules = $4;
				}
				;

modules: { $$ = []string{}; /* Must return an empty type. */}
			 | modules module { $$ = append($$, $2); }
			 ;
