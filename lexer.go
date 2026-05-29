package main

import (
	"unicode"
)

const (
	INVALID_TOKEN = iota
	EOF_TOKEN

	DECLARE_TOKEN
	OF_TOKEN
	DEFINE_TOKEN
	REPEAT_TOKEN
	UNTIL_TOKEN
	PROCEDURE_TOKEN
	END_PROCEDURE_TOKEN
	FUNCTION_TOKEN
	END_FUNCTION_TOKEN
	RETURNS_TOKEN
	RETURN_TOKEN
	FOR_TOKEN
	TO_TOKEN
	STEP_TOKEN
	NEXT_TOKEN
	TYPE_TOKEN
	END_TYPE_TOKEN
	CASE_TOKEN
	END_CASE_TOKEN
	OTHERWISE_TOKEN
	CALL_TOKEN
	WHILE_TOKEN
	END_WHILE_TOKEN
	BYREF_TOKEN
	IF_TOKEN
	THEN_TOKEN
	ELSE_TOKEN
	END_IF_TOKEN

	// Quality of life
	BREAK_TOKEN
	CONTINUE_TOKEN

	// OOP
	CLASS_TOKEN
	INHERITS_TOKEN
	NEW_TOKEN
	END_CLASS_TOKEN
	PUBLIC_TOKEN
	PRIVATE_TOKEN

	// Symbol Tokens
	LEFT_ARROW_TOKEN
	COMMA_TOKEN
	FULL_STOP_TOKEN
	COLON_TOKEN
	CARET_TOKEN

	// Arithmetic Operators
	ADD_TOKEN
	SUBTRACT_TOKEN
	MULTIPLICATION_TOKEN
	DIVISION_TOKEN
	DIV_TOKEN
	MOD_TOKEN
	BIT_AND_TOKEN

	// Bool
	AND_TOKEN
	OR_TOKEN
	NOT_TOKEN

	// Comparison Operators
	GREATER_TOKEN
	GREATER_EQ_TOKEN
	LESSER_TOKEN
	LESSER_EQ_TOKEN
	EQUAL_TOKEN
	NOT_EQUAL_TOKEN

	// Brackets
	OPEN_SQUARE_TOKEN
	CLOSE_SQUARE_TOKEN
	OPEN_CURLY_TOKEN
	CLOSE_CURLY_TOKEN
	OPEN_PAREN_TOKEN
	CLOSE_PAREN_TOKEN

	// Value Tokens
	IDENTIFIER_TOKEN
	STRING_TOKEN
	INT_TOKEN
	REAL_TOKEN
	CHAR_TOKEN
	DATE_TOKEN

	// Builtins
	OUTPUT_TOKEN
	INPUT_TOKEN
	OPENFILE_TOKEN
	SEEK_TOKEN
	CLOSEFILE_TOKEN
	READFILE_TOKEN
	WRITEFILE_TOKEN
	FILE_EOF_TOKEN
)

var TOKENS = map[string]uint8{
	"DECLARE":      DECLARE_TOKEN,
	"OF":           OF_TOKEN,
	"DEFINE":       DEFINE_TOKEN,
	"REPEAT":       REPEAT_TOKEN,
	"UNTIL":        UNTIL_TOKEN,
	"PROCEDURE":    PROCEDURE_TOKEN,
	"ENDPROCEDURE": END_PROCEDURE_TOKEN,
	"FUNCTION":     FUNCTION_TOKEN,
	"ENDFUNCTION":  END_FUNCTION_TOKEN,
	"RETURNS":      RETURNS_TOKEN,
	"RETURN":       RETURN_TOKEN,
	"FOR":          FOR_TOKEN,
	"TO":           TO_TOKEN,
	"STEP":         STEP_TOKEN,
	"NEXT":         NEXT_TOKEN,
	"TYPE":         TYPE_TOKEN,
	"ENDTYPE":      END_TYPE_TOKEN,
	"CASE":         CASE_TOKEN,
	"ENDCASE":      END_CASE_TOKEN,
	"OTHERWISE":    OTHERWISE_TOKEN,
	"CALL":         CALL_TOKEN,
	"WHILE":        WHILE_TOKEN,
	"ENDWHILE":     END_WHILE_TOKEN,
	"BYREF":        BYREF_TOKEN,
	"IF":           IF_TOKEN,
	"THEN":         THEN_TOKEN,
	"ELSE":         ELSE_TOKEN,
	"ENDIF":        END_IF_TOKEN,

	// Quality of life
	"BREAK":    BREAK_TOKEN,
	"CONTINUE": CONTINUE_TOKEN,

	// OOP
	"CLASS":    CLASS_TOKEN,
	"INHERITS": INHERITS_TOKEN,
	"NEW":      NEW_TOKEN,
	"ENDCLASS": END_CLASS_TOKEN,
	"PUBLIC":   PUBLIC_TOKEN,
	"PRIVATE":  PRIVATE_TOKEN,

	// Symbol Tokens
	"<-": LEFT_ARROW_TOKEN,
	",":  COMMA_TOKEN,
	".":  FULL_STOP_TOKEN,
	":":  COLON_TOKEN,
	"^":  CARET_TOKEN,

	// Arithmetic Operators
	"+":   ADD_TOKEN,
	"-":   SUBTRACT_TOKEN,
	"*":   MULTIPLICATION_TOKEN,
	"/":   DIVISION_TOKEN,
	"DIV": DIV_TOKEN,
	"MOD": MOD_TOKEN,
	"&":   BIT_AND_TOKEN,

	// Bool
	"AND": AND_TOKEN,
	"OR":  OR_TOKEN,
	"NOT": NOT_TOKEN,

	// Comparison Operators
	">":  GREATER_TOKEN,
	">=": GREATER_EQ_TOKEN,
	"<":  LESSER_TOKEN,
	"<=": LESSER_EQ_TOKEN,
	"=":  EQUAL_TOKEN,
	"<>": NOT_EQUAL_TOKEN,

	// Brackets
	"[": OPEN_SQUARE_TOKEN,
	"]": CLOSE_SQUARE_TOKEN,
	"{": OPEN_CURLY_TOKEN,
	"}": CLOSE_CURLY_TOKEN,
	"(": OPEN_PAREN_TOKEN,
	")": CLOSE_PAREN_TOKEN,

	// Builtins
	"OUTPUT":    OUTPUT_TOKEN,
	"INPUT":     INPUT_TOKEN,
	"OPENFILE":  OPENFILE_TOKEN,
	"SEEK":      SEEK_TOKEN,
	"CLOSEFILE": CLOSEFILE_TOKEN,
	"READFILE":  READFILE_TOKEN,
	"WRITEFILE": WRITEFILE_TOKEN,
	"EOF":       FILE_EOF_TOKEN,
}

type Pos struct {
	Line uint16
	Col  uint16
}

type Span struct {
	From Pos
	To   Pos
}

type Token struct {
	TokenType uint8
	Value     string
	Span      Span
}

func tokenize(text string) []Token {
	index := 0
	tokens := make([]Token, 0, 256)
	pos := Pos{
		Line: 1,
		Col:  1,
	}

	symbols := make([]byte, 0, 16)

	for index < len(text) {
		c := text[index]

		if c == '\n' {
			symbols = symbols[:0]
			pos.Line += 1
			pos.Col = 1
			continue
		} else if unicode.IsSpace(rune(c)) {
			symbols = symbols[:0]
			pos.Col += 1
			index += 1
			continue
		} else if unicode.IsDigit(rune(c)) {
			symbols = symbols[:0]
			num := ""
			is_real := false
			from := pos

			for unicode.IsDigit(rune(c)) || c == '_' || c == '.' {
				if c == '.' {
					is_real = true
					num += string(c)
				} else if c != '_' {
					num += string(c)
				}

				pos.Col += 1
				index += 1
				c = text[index]
			}

			if is_real {
				tokens = append(tokens, Token{
					TokenType: REAL_TOKEN,
					Value:     num,
					Span:      Span{From: from, To: pos},
				})
			} else {
				tokens = append(tokens, Token{
					TokenType: INT_TOKEN,
					Value:     num,
					Span:      Span{From: from, To: pos},
				})
			}
		} else if unicode.IsLetter(rune(c)) || c == '_' {
			symbols = symbols[:0]
			value := ""
			from := pos

			for unicode.IsLetter(rune(c)) || unicode.IsDigit(rune(c)) || c == '_' {
				value += string(c)

				pos.Col += 1
				index += 1
				c = text[index]
			}

			map_token := TOKENS[value]
			if map_token != INVALID_TOKEN {
				tokens = append(tokens, Token{
					TokenType: map_token,
					Value:     value,
					Span:      Span{From: from, To: pos},
				})
			} else {
				tokens = append(tokens, Token{
					TokenType: IDENTIFIER_TOKEN,
					Value:     value,
					Span:      Span{From: from, To: pos},
				})
			}
		} else if c == '"' || c == '\'' {
			symbols = symbols[:0]
			value := ""
			from := pos
			is_char := c == '\''

			for (is_char && c != '\'') || (!is_char && c != '"') {
				value += string(c)
				if c == '\n' {
					pos.Line += 1
					pos.Col = 1
				} else {
					pos.Col += 1
				}

				index += 1
				c = text[index]
			}

			if is_char {
				tokens = append(tokens, Token{
					TokenType: CHAR_TOKEN,
					Value:     value,
					Span:      Span{From: from, To: pos},
				})
			} else {
				tokens = append(tokens, Token{
					TokenType: STRING_TOKEN,
					Value:     value,
					Span:      Span{From: from, To: pos},
				})
			}
		} else {
			symbols = append(symbols, c)
			value := string(symbols)
			map_token := TOKENS[value]
			if map_token != INVALID_TOKEN {
				tokens = append(tokens, Token{
					TokenType: map_token,
					Value:     value,
					Span: Span{From: Pos{
						Line: pos.Line,
						Col:  pos.Col - uint16(len(symbols)),
					}, To: pos},
				})

				symbols = symbols[:0]
			}

			index += 1
		}
	}

	tokens = append(tokens, Token{
		TokenType: EOF_TOKEN,
		Value:     "EOF",
		Span:      Span{},
	})

	return tokens
}
