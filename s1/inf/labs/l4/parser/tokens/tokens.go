package tokens

type TokenType string

const (
	TokenTypeIdentifier     TokenType = "identifier"
	TokenTypeString         TokenType = "string"
	TokenTypeNumber         TokenType = "number"
	TokenTypeBinaryOperator TokenType = "binary_operator"
	TokenTypeUnaryOperator  TokenType = "unary_operator"
	TokenTypeLeftBrace      TokenType = "left_brace"
	TokenTypeRightBrace     TokenType = "right_brace"
	TokenTypeLeftBracket    TokenType = "left_bracket"
	TokenTypeRightBracket   TokenType = "right_bracket"
	TokenTypeLeftParen      TokenType = "left_paren"
	TokenTypeRightParen     TokenType = "right_paren"
	TokenTypeAssign         TokenType = "assign"
	TokenTypeComma          TokenType = "comma"
	TokenTypeTrue           TokenType = "true"
	TokenTypeFalse          TokenType = "false"
	TokenTypeNull           TokenType = "null"
	TokenTypeNewLine        TokenType = "new_line"
	TokenTypeComment        TokenType = "comment"
	TokenTypeEOF            TokenType = "EOF"
)

type Token struct {
	value     string
	tokenType TokenType
	metadata  TokenMetadata
}

type TokenMetadata struct {
	Line     int
	Position int
}

func NewToken(tokenType TokenType, value string) *Token {
	return &Token{
		value:     value,
		tokenType: tokenType,
		metadata:  TokenMetadata{},
	}
}

func (t *Token) Type() TokenType {
	return t.tokenType
}

func (t *Token) Value() string {
	return t.value
}

func (t *Token) Metadata() TokenMetadata {
	return t.metadata
}

func (t *Token) SetMetadata(m TokenMetadata) *Token {
	t.metadata = m
	return t
}

func NewIdentifierToken(value string) *Token {
	return NewToken(TokenTypeIdentifier, value)
}

func NewStringToken(value string) *Token {
	return NewToken(TokenTypeString, value)
}

func NewNumberToken(value string) *Token {
	return NewToken(TokenTypeNumber, value)
}

func NewBinaryOperator(value string) *Token {
	return NewToken(TokenTypeBinaryOperator, value)
}

func NewUnaryOperator(value string) *Token {
	return NewToken(TokenTypeUnaryOperator, value)
}

func NewCommentToken(value string) *Token {
	return NewToken(TokenTypeComment, value)
}

var (
	LeftBrace    = NewToken(TokenTypeLeftBrace, "{")
	RightBrace   = NewToken(TokenTypeRightBrace, "}")
	LeftBracket  = NewToken(TokenTypeLeftBracket, "[")
	RightBracket = NewToken(TokenTypeRightBracket, "]")
	LeftParen    = NewToken(TokenTypeLeftParen, "(")
	RightParen   = NewToken(TokenTypeRightParen, ")")
	True         = NewToken(TokenTypeTrue, "true")
	False        = NewToken(TokenTypeFalse, "false")
	Null         = NewToken(TokenTypeNull, "null")
	Comma        = NewToken(TokenTypeComma, ",")
	Assign       = NewToken(TokenTypeAssign, "")
	NewLine      = NewToken(TokenTypeNewLine, "")
	EOF          = NewToken(TokenTypeEOF, "")
)

// package tokens
//

// type TokenType string

// const (
// 	TokenTypeBracket TokenType = "bracket"

// 	TokenTypeName   TokenType = "name"
// 	TokenTypeString TokenType = "string"
// 	TokenTypeNumber TokenType = "number"
// 	TokenTypeValue  TokenType = "value"

// 	TokenTypeOperator TokenType = "operator"

// 	TokenTypeAssign TokenType = "assign"
// 	TokenTypeComma  TokenType = "comma"

// 	TokenTypeComment TokenType = "comment"

// 	TokenTypeNewLine TokenType = "new_line"
// )

// type Token interface {
// 	Type() TokenType
// }

// type BracketToken struct {
// 	char rune
// }

// func NewBracketToken(char rune) *BracketToken {
// 	return &BracketToken{
// 		char: char,
// 	}
// }

// func (b *BracketToken) Type() TokenType {
// 	return TokenTypeBracket
// }

// func (b *BracketToken) String() string {
// 	return string(b.char)
// }

// type StringToken struct {
// 	str string
// }

// func NewStringToken(str string) *StringToken {
// 	return &StringToken{
// 		str: str,
// 	}
// }

// func (s *StringToken) Type() TokenType {
// 	return TokenTypeString
// }

// func (s *StringToken) String() string {
// 	return s.str
// }

// type NumberToken struct {
// 	str string
// }

// func NewNumberToken(str string) *NumberToken {
// 	return &NumberToken{
// 		str: str,
// 	}
// }

// func (n *NumberToken) Type() TokenType {
// 	return TokenTypeNumber
// }

// func (n *NumberToken) String() string {
// 	return n.str
// }

// type ValueToken struct {
// 	str string
// }

// func NewValueToken(str string) *ValueToken {
// 	return &ValueToken{
// 		str: str,
// 	}
// }

// func (v *ValueToken) Type() TokenType {
// 	return TokenTypeValue
// }

// func (v *ValueToken) String() string {
// 	return v.str
// }

// type OperatorToken struct {
// 	str string
// }

// func NewOperatorToken(str string) *OperatorToken {
// 	return &OperatorToken{
// 		str: str,
// 	}
// }

// func (n *OperatorToken) Type() TokenType {
// 	return TokenTypeOperator
// }

// func (n *OperatorToken) String() string {
// 	return n.str
// }

// type NameToken struct {
// 	str string
// }

// func NewNameToken(str string) *NameToken {
// 	return &NameToken{
// 		str: str,
// 	}
// }

// func (n *NameToken) Type() TokenType {
// 	return TokenTypeName
// }

// func (n *NameToken) String() string {
// 	return n.str
// }

// type AssignToken struct{}

// func NewAssignToken() *AssignToken {
// 	return &AssignToken{}
// }

// func (a *AssignToken) Type() TokenType {
// 	return TokenTypeAssign
// }

// type NewLineToken struct{}

// func NewNewLineToken() *NewLineToken {
// 	return &NewLineToken{}
// }

// func (n *NewLineToken) Type() TokenType {
// 	return TokenTypeNewLine
// }

// type CommaToken struct{}

// func NewCommaToken() *CommaToken {
// 	return &CommaToken{}
// }

// func (c *CommaToken) Type() TokenType {
// 	return TokenTypeComma
// }

// type CommentToken struct {
// 	str string
// }

// func NewCommentToken(str string) *CommentToken {
// 	return &CommentToken{str: str}
// }

// func (c *CommentToken) Type() TokenType {
// 	return TokenTypeComment
// }

// func (c *CommentToken) String() string {
// 	return c.str
// }
