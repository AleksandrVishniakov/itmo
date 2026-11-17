package hcl

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"iter"
	"parser/tokens"
	"slices"
	"strings"
	"unicode"
)

const (
	leftBracket  rune = '['
	rightBracket rune = ']'
	leftBrace    rune = '{'
	rightBrace   rune = '}'
	leftParen    rune = '('
	rightParen   rune = ')'

	doubleQuote rune = '"'

	dot       rune = '.'
	comma     rune = ','
	dash      rune = '-'
	underline rune = '_'

	plus     rune = '+'
	minus    rune = '-'
	multiply rune = '*'
	divide   rune = '/'

	greaterThan rune = '>'
	lessThan    rune = '<'
	equals      rune = '='

	exclamationMark rune = '!'
	pipe            rune = '|'
	ampersand       rune = '&'

	assign rune = '='

	hash      rune = '#'
	slash     rune = '/'
	backslash rune = '\\'
	asterisk  rune = '*'

	newline rune = '\n'
)

var (
	operators = []string{"<=", ">=", "==", "!=", "&&", "||",
		"+", "-", "*", "/", "=", "!", ">", "<"}
	unaryOperators = []string{"!"}
)

type LexerError struct {
	Line   int
	Column int
	Err    error
}

func (e *LexerError) Error() string {
	return fmt.Sprintf("lexer error near line %d column %d: %s", e.Line, e.Column, e.Err)
}

func (e *LexerError) Unwrap() error {
	return e.Err
}

type Lexer struct {
	reader *bufio.Reader
	line   int
	column int
}

func NewLexer(r io.Reader) *Lexer {
	return &Lexer{
		reader: bufio.NewReader(r),
		line:   1,
		column: 1,
	}
}

func (l *Lexer) Tokens() iter.Seq2[*tokens.Token, error] {
	return func(yield func(*tokens.Token, error) bool) {
		for {
			token, err := l.NextToken()

			if !yield(token, err) {
				return
			}

			if token.Type() == tokens.TokenTypeEOF {
				return
			}
		}
	}
}

func (l *Lexer) NextToken() (token *tokens.Token, err error) {
	defer func() {
		if token != nil {
			token = token.SetMetadata(tokens.TokenMetadata{
				Line:     l.line,
				Position: l.column,
			})
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				err = nil
				token = tokens.EOF
			}
		}
	}()

	if err := l.SkipWhitespaces(); err != nil {
		if errors.Is(err, io.EOF) {
			return tokens.EOF, nil
		}

		return nil, l.WrapError(err)
	}

	next, err := l.PeekRune()
	if err != nil {
		return nil, l.WrapError(err)
	}

	startsWithHash, err := l.StartsWith("#")
	if err != nil {
		return nil, l.WrapError(err)
	}

	startsWithSlashes, err := l.StartsWith("//")
	if err != nil {
		return nil, l.WrapError(err)
	}

	startsMultiLineComment, err := l.StartsWith("/*")
	if err != nil {
		return nil, l.WrapError(err)
	}

	isHeredoc, err := l.StartsWith("<<")
	if err != nil {
		return nil, err
	}

	isOperator, err := l.isOperator()
	if err != nil {
		return nil, l.WrapError(err)
	}

	switch {
	case next == newline:
		token, err = l.tokenizeNewLine()
		if err != nil {
			return nil, l.WrapError(err)
		}
		return token, nil
	case next == comma:
		token, err = l.tokenizeComma()
		if err != nil {
			return nil, l.WrapError(err)
		}
		return token, nil
	case next == doubleQuote:
		token, err = l.tokenizeString()
		if err != nil {
			return nil, l.WrapError(err)
		}
		return token, nil
	case isBracket(next):
		token, err = l.tokenizeBracket()
		if err != nil {
			return nil, l.WrapError(err)
		}
		return token, nil
	case unicode.IsLetter(next) || next == underline:
		token, err = l.tokenizeIdentifier()
		if err != nil {
			return nil, l.WrapError(err)
		}
		return token, nil
	case unicode.IsDigit(next):
		token, err = l.tokenizeNumber()
		if err != nil {
			return nil, l.WrapError(err)
		}
		return token, nil
	case startsWithHash || startsWithSlashes:
		token, err = l.tokenizeSingleLineComment()
		if err != nil {
			return nil, l.WrapError(err)
		}
		return token, nil
	case startsMultiLineComment:
		token, err = l.tokenizeMultiLineComment()
		if err != nil {
			return nil, l.WrapError(err)
		}
		return token, nil
	case isHeredoc:
		token, err = l.tokenizeHeredoc()
		if err != nil {
			return nil, l.WrapError(err)
		}
		return token, nil
	case isOperator:
		token, err = l.tokenizeOperator()
		if err != nil {
			return nil, l.WrapError(err)
		}
		return token, nil
	}

	return nil, l.Error(fmt.Sprintf("unexpected symbol: %q", next))
}

func (l *Lexer) tokenizeHeredoc() (*tokens.Token, error) {
	// skip "<<"
	for range 2 {
		_, err := l.NextRune()
		if err != nil {
			return nil, err
		}
	}

	next, err := l.PeekRune()
	if err != nil {
		return nil, err
	}

	removeIndent := next == dash
	if removeIndent {
		_, err := l.NextRune()
		if err != nil {
			return nil, err
		}
	}

	if err := l.SkipWhitespaces(); err != nil {
		return nil, err
	}

	marker, err := l.tokenizeIdentifier()
	if err != nil {
		return nil, err
	}

	if strings.Contains(marker.Value(), string(dot)) {
		return nil, fmt.Errorf("invalid heredoc marker: %q", marker.Value())
	}

	next, err = l.NextRune()
	if err != nil {
		return nil, err
	}

	if next != newline {
		return nil, fmt.Errorf("invalid heredoc marker: %q", marker.Value())
	}

	lines := make([]string, 0)
	builder := strings.Builder{}
	for {
		end, err := l.StartsWith(marker.Value())
		if err != nil {
			return nil, err
		}

		if end && strings.TrimSpace(builder.String()) == "" {
			break
		}

		next, err = l.PeekRune()
		if err != nil {
			return nil, err
		}

		if next == newline {
			lines = append(lines, builder.String())
			builder.Reset()
		} else {
			builder.WriteRune(next)
		}

		_, err = l.NextRune()
		if err != nil {
			return nil, err
		}
	}

	for range len(marker.Value()) {
		_, err = l.NextRune()
		if err != nil {
			return nil, err
		}
	}

	if removeIndent {
		lines = removeCommonIndent(lines)
	}

	return tokens.NewStringToken(strings.Join(lines, string(newline))), nil
}

func removeCommonIndent(lines []string) []string {
	minIndent := -1
	for _, line := range lines {
		cnt := 0
		for _, ch := range line {
			if ch == ' ' || ch == '\t' {
				cnt += 1
			} else {
				break
			}
		}

		if minIndent == -1 {
			minIndent = cnt
		} else {
			minIndent = min(minIndent, cnt)
		}
	}

	for i := range lines {
		lines[i] = lines[i][minIndent:]
	}

	return lines
}

func (l *Lexer) tokenizeOperator() (*tokens.Token, error) {
	for _, op := range operators {
		ok, err := l.StartsWith(op)
		if err != nil {
			return nil, err
		}

		if ok {
			for range len(op) {
				_, err = l.NextRune()
				if err != nil {
					return nil, err
				}
			}

			if op == string(assign) {
				return tokens.Assign, nil
			}

			if slices.Contains(unaryOperators, op) {
				return tokens.NewUnaryOperator(op), nil
			} else {
				return tokens.NewBinaryOperator(op), nil
			}
		}
	}

	return nil, errors.New("unexpected operator")
}

func (l *Lexer) tokenizeMultiLineComment() (*tokens.Token, error) {
	// skip "/*"
	for range 2 {
		_, err := l.NextRune()
		if err != nil {
			return nil, err
		}
	}

	builder := strings.Builder{}
	for {
		end, err := l.StartsWith("*/")
		if err != nil {
			return nil, err
		}

		if end {
			// skip "*/"
			for range 2 {
				_, err := l.NextRune()
				if err != nil {
					return nil, err
				}
			}
			return tokens.NewCommentToken(builder.String()), nil
		}

		ch, err := l.NextRune()
		if err != nil {
			return nil, err
		}

		builder.WriteRune(ch)
	}
}

func (l *Lexer) tokenizeSingleLineComment() (*tokens.Token, error) {
	startsWithHash, err := l.StartsWith("#")
	if err != nil {
		return nil, err
	}

	startsWithSlashes, err := l.StartsWith("//")
	if err != nil {
		return nil, err
	}

	switch {
	case startsWithHash:
		_, err := l.NextRune()
		if err != nil {
			return nil, err
		}
	case startsWithSlashes:
		for range 2 {
			_, err := l.NextRune()
			if err != nil {
				return nil, err
			}
		}
	default:
		return nil, errors.New("unexpected comment type")
	}

	if err := l.SkipWhitespaces(); err != nil {
		return nil, err
	}

	builder := strings.Builder{}
	for {
		next, err := l.PeekRune()
		if err != nil {
			return nil, err
		}

		if next == newline {
			return tokens.NewCommentToken(builder.String()), nil
		}

		builder.WriteRune(next)

		_, err = l.NextRune()
		if err != nil {
			return nil, err
		}
	}
}

func (l *Lexer) tokenizeNumber() (*tokens.Token, error) {
	var prev rune
	builder := strings.Builder{}
	var wasE, wasDot bool

	for {
		ch, err := l.PeekRune()
		if err != nil {
			return nil, err
		}

		if unicode.IsDigit(ch) {
			builder.WriteRune(ch)
		} else if ch == 'e' || ch == 'E' {
			if wasE || !unicode.IsDigit(prev) {
				return nil, errors.New("invalid number notation")
			}
			wasE = true
			builder.WriteRune(ch)
		} else if ch == dot {
			if wasDot || wasE {
				return nil, errors.New("invalid number notation")
			}
			wasDot = true
			builder.WriteRune(ch)
		} else if ch == minus || ch == plus {
			if prev != 'e' && prev != 'E' {
				break
			}
			builder.WriteRune(ch)
		} else {
			break
		}

		prev = ch
		_, err = l.NextRune()
		if err != nil {
			return nil, err
		}
	}

	if !unicode.IsDigit(prev) {
		return nil, errors.New("invalid number notation")
	}

	return tokens.NewNumberToken(builder.String()), nil
}

func (l *Lexer) tokenizeIdentifier() (*tokens.Token, error) {
	var prev rune
	builder := strings.Builder{}

	for {
		ch, err := l.PeekRune()
		if err != nil {
			return nil, err
		}

		if unicode.IsLetter(ch) || ch == underline {
			builder.WriteRune(ch)
		} else if unicode.IsDigit(ch) || ch == dash || ch == dot {
			if prev == dot {
				return nil, errors.New("invalid name")
			}

			builder.WriteRune(ch)
		} else {
			break
		}

		prev = ch
		_, err = l.NextRune()
		if err != nil {
			return nil, err
		}
	}

	switch builder.String() {
	case "true":
		return tokens.True, nil
	case "false":
		return tokens.False, nil
	case "null":
		return tokens.Null, nil
	default:
		return tokens.NewIdentifierToken(builder.String()), nil
	}
}

func (l *Lexer) tokenizeBracket() (*tokens.Token, error) {
	ch, err := l.NextRune()
	if err != nil {
		return nil, err
	}

	switch ch {
	case leftBrace:
		return tokens.LeftBrace, nil
	case rightBrace:
		return tokens.RightBrace, nil
	case leftBracket:
		return tokens.LeftBracket, nil
	case rightBracket:
		return tokens.RightBracket, nil
	case leftParen:
		return tokens.LeftParen, nil
	case rightParen:
		return tokens.RightParen, nil
	default:
		return nil, fmt.Errorf("unexpected symbol: %q", ch)
	}
}

func (l *Lexer) tokenizeNewLine() (*tokens.Token, error) {
	_, err := l.NextRune()
	if err != nil {
		return nil, err
	}

	return tokens.NewLine, nil
}

func (l *Lexer) tokenizeComma() (*tokens.Token, error) {
	_, err := l.NextRune()
	if err != nil {
		return nil, err
	}

	return tokens.Comma, nil
}

func (l *Lexer) tokenizeString() (*tokens.Token, error) {
	ch, err := l.NextRune()
	if err != nil {
		return nil, err
	}

	if ch != doubleQuote {
		return nil, fmt.Errorf("invalid string start: %q", ch)
	}

	escaped := false
	builder := strings.Builder{}
	for {
		ch, err := l.NextRune()
		if err != nil {
			return nil, err
		}

		switch ch {
		case doubleQuote:
			if escaped {
				builder.WriteRune(ch)
			} else {
				return tokens.NewStringToken(builder.String()), nil
			}
		case newline:
			if escaped {
				builder.WriteRune(ch)
			} else {
				return nil, errors.New("unexpected string end")
			}
		case backslash:
			if escaped {
				builder.WriteRune(ch)
			}
		default:
			builder.WriteRune(ch)
		}

		escaped = ch == backslash
	}
}

func (l *Lexer) isOperator() (bool, error) {
	for _, op := range operators {
		ok, err := l.StartsWith(op)
		if err != nil {
			return false, err
		}

		if ok {
			return true, nil
		}
	}
	return false, nil
}

func (l *Lexer) NextRune() (rune, error) {
	r, _, err := l.reader.ReadRune()
	if err != nil {
		return 0, err
	}

	if r == '\n' {
		l.column = 1
		l.line += 1
	} else {
		l.column += 1
	}

	return r, nil
}

func (l *Lexer) PeekRune() (rune, error) {
	r, _, err := l.reader.ReadRune()
	if err != nil {
		return 0, err
	}

	err = l.reader.UnreadRune()
	if err != nil {
		return 0, err
	}

	return r, nil
}

func (l *Lexer) SkipWhitespaces() error {
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			return err
		}

		if r != ' ' && r != '\t' {
			err = l.reader.UnreadRune()
			if err != nil {
				return err
			}

			return nil
		} else {
			l.column += 1
		}
	}
}

func (l *Lexer) StartsWith(text string) (bool, error) {
	b, err := l.reader.Peek(len(text))
	if err != nil {
		return false, err
	}

	return string(b) == text, nil
}

func (l *Lexer) TokenMetadata() tokens.TokenMetadata {
	return tokens.TokenMetadata{
		Line:     l.line,
		Position: l.column,
	}
}

func (l *Lexer) Error(text string) error {
	return &LexerError{
		Line:   l.line,
		Column: l.column,
		Err:    fmt.Errorf("hcl.Lexer: %s", text),
	}
}

func (l *Lexer) WrapError(err error) error {
	return &LexerError{
		Line:   l.line,
		Column: l.column,
		Err:    fmt.Errorf("hcl.Lexer: %w", err),
	}
}

func isBracket(r rune) bool {
	return r == leftBracket || r == rightBracket ||
		r == leftBrace || r == rightBrace ||
		r == leftParen || r == rightParen
}
