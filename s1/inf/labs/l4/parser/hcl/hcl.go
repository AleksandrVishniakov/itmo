package hcl

import (
	"io"
	"parser/nodes"
	"parser/tokens"
)

type HCL struct {
}

func (h HCL) ToAST(r io.Reader) (nodes.Value, error) {
	lexer := NewLexer(r)

	t := make([]*tokens.Token, 0)
	for token, err := range lexer.Tokens() {
		if err != nil {
			return nil, err
		}

		if token.Type() != tokens.TokenTypeComment {
			t = append(t, token)
		}
	}

	parser := NewParser(t)
	file, err := parser.Parse()
	if err != nil {
		return nil, err
	}

	ast, err := parser.Convert(file)
	if err != nil {
		return nil, err
	}

	return ast, nil
}
