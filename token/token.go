package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// 識別子 + リテラル
	ILLEGAL  = "ILLEGAL"
	EOF      = "EOF"
	IDENT    = "IDENT"
	INT      = "INT"
	MINUS    = "MINUS"
	BANG     = "BANG"
	ASTERISK = "ASTERISK"
	SLASH    = "SLASH"

	LT = "<"
	GT = ">"

	// 演算子
	ASSIGN = "ASSIGN"
	PLUS   = "PLUS"
	EQ     = "=="
	NOT_EQ = "!="

	// デリミタ
	COMMA     = "COMMA"
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIndent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
