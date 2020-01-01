package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "IILLEGAL" // 未知の文字
	EOF     = "EOF"      // ファイル終端

	// 識別子, リテラル
	IDENT  = "IDENT"  // ユーザ定義識別子
	INT    = "INT"    // 数値リテラル
	STRING = "STRING" // 文字列リテラル

	// 演算子
	// 代入演算子
	ASSIGN = "="
	// 算術演算子
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	// 比較演算子
	EQ     = "=="
	NOT_EQ = "!="
	LT     = "<"
	GT     = ">"

	// デリミタ
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

// キーワード
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// 与えられた文字列が予約されている場合はそのTokenTypeを返し、
// それ以外はIDENTを返す
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
