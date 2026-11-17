package parser

import (
	"io"
	"parser/nodes"
)

type ASTGenerator interface {
	ToAST(r io.Reader) (nodes.Value, error)
}

type ASTResolver interface {
	FromAST(io.Writer, nodes.Value) error
}

type Parser struct {
	from ASTGenerator
	to   ASTResolver
}

func NewParser(from ASTGenerator, to ASTResolver) *Parser {
	return &Parser{
		from: from,
		to:   to,
	}
}

func (p *Parser) Parse(r io.Reader, w io.Writer) error {
	ast, err := p.from.ToAST(r)
	if err != nil {
		return err
	}

	return p.to.FromAST(w, ast)
}
