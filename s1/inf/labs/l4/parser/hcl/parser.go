package hcl

import (
	"fmt"
	"parser/nodes"
	"parser/tokens"
	"slices"
	"strings"
)

type Parser struct {
	tokens   []*tokens.Token
	position int
}

func NewParser(tokens []*tokens.Token) *Parser {
	return &Parser{
		tokens:   tokens,
		position: 0,
	}
}

func (p *Parser) Parse() (*File, error) {
	return p.parseFile()
}

func (p *Parser) current() *tokens.Token {
	return p.tokens[p.position]
}

func (p *Parser) peek(n int) *tokens.Token {
	if p.position+n >= len(p.tokens) {
		return p.tokens[len(p.tokens)-1]
	}

	return p.tokens[p.position+n]
}

func (p *Parser) nextToken() *tokens.Token {
	p.position += 1
	return p.tokens[p.position]
}

func (p *Parser) match(expected tokens.TokenType) bool {
	if p.current().Type() == expected {
		p.nextToken()
		return true
	}

	return false
}

func (p *Parser) expect(expected tokens.TokenType) (*tokens.Token, error) {
	if token := p.current(); token.Type() == expected {
		p.nextToken()
		return token, nil
	}

	return nil, p.wrapError(fmt.Errorf("unexpected token: expected %q, got %q", expected, p.current().Type()))
}

func (p *Parser) skipNewLines() {
	for p.current().Type() == tokens.TokenTypeNewLine {
		p.nextToken()
	}
}

func (p *Parser) wrapError(err error) error {
	t := p.current()

	return fmt.Errorf("hcl.Parser: error near line %d column %d %q: %w", t.Metadata().Line, t.Metadata().Position, t.Value(), err)
}

func (p *Parser) parseString() (*String, error) {
	token, err := p.expect(tokens.TokenTypeString)
	if err != nil {
		return nil, err
	}

	return new(String(token.Value())), nil
}

func (p *Parser) parseNumber() (*Number, error) {
	token, err := p.expect(tokens.TokenTypeNumber)
	if err != nil {
		return nil, err
	}

	return new(Number(token.Value())), nil
}

func (p *Parser) parseBoolean() (*Boolean, error) {
	switch p.current().Type() {
	case tokens.TokenTypeTrue:
		p.nextToken()
		return new(Boolean(true)), nil
	case tokens.TokenTypeFalse:
		p.nextToken()
		return new(Boolean(false)), nil
	default:
		return nil, p.wrapError(fmt.Errorf("unexpected token"))
	}
}

func (p *Parser) parseNull() (*Null, error) {
	_, err := p.expect(tokens.TokenTypeNull)
	if err != nil {
		return nil, err
	}

	return new(Null), nil
}

func (p *Parser) parseArray() (*Array, error) {
	if _, err := p.expect(tokens.TokenTypeLeftBracket); err != nil {
		return nil, err
	}

	elements := make([]Node, 0)
	p.skipNewLines()

	if p.match(tokens.TokenTypeRightBracket) {
		return new(Array(elements)), nil
	}

	for {
		p.skipNewLines()

		value, err := p.parseExpression()
		if err != nil {
			return nil, err
		}

		elements = append(elements, value)
		p.skipNewLines()

		if p.match(tokens.TokenTypeRightBracket) {
			break
		}

		if _, err := p.expect(tokens.TokenTypeComma); err != nil {
			return nil, err
		}
	}

	return new(Array(elements)), nil
}

func (p *Parser) parseObject() (*Object, error) {
	if _, err := p.expect(tokens.TokenTypeLeftBrace); err != nil {
		return nil, err
	}

	var attributes []Node

	p.skipNewLines()

	if p.match(tokens.TokenTypeRightBrace) {
		return new(Object(attributes)), nil
	}

	for {
		key, err := p.expect(tokens.TokenTypeIdentifier)
		if err != nil {
			return nil, err
		}

		if _, err := p.expect(tokens.TokenTypeAssign); err != nil {
			return nil, err
		}

		value, err := p.parseExpression()
		if err != nil {
			return nil, err
		}

		attributes = append(attributes, &Attribute{
			Name:  key.Value(),
			Value: value,
		})

		hadNewline := p.match(tokens.TokenTypeNewLine)
		p.skipNewLines()

		if p.match(tokens.TokenTypeRightBrace) {
			break
		}

		if p.match(tokens.TokenTypeComma) {
			p.skipNewLines()
			continue
		}

		if hadNewline {
			continue
		}

		return nil, p.wrapError(fmt.Errorf("unexpected token"))
	}

	return new(Object(attributes)), nil
}

func (p *Parser) parseExpression() (Node, error) {
	switch p.current().Type() {
	case tokens.TokenTypeLeftBracket:
		return p.parseArray()
	case tokens.TokenTypeLeftBrace:
		return p.parseObject()
	}

	if p.current().Value() == "-" &&
		p.peek(1).Type() == tokens.TokenTypeNumber {
		p.nextToken()
		number, err := p.expect(tokens.TokenTypeNumber)
		if err != nil {
			return nil, err
		}

		return new(Number("-" + number.Value())), nil
	}

	cnt := 0
	for p.peek(cnt).Type() != tokens.TokenTypeNewLine &&
		p.peek(cnt).Type() != tokens.TokenTypeEOF &&
		p.peek(cnt).Type() != tokens.TokenTypeComma &&
		p.peek(cnt).Type() != tokens.TokenTypeRightBrace &&
		p.peek(cnt).Type() != tokens.TokenTypeRightBracket {
		cnt++
	}

	if cnt == 0 {
		return nil, p.wrapError(fmt.Errorf("unexpected token"))
	}

	if cnt > 1 {
		parts := []string{p.current().Value()}

		for i := 1; i < cnt; i++ {
			parts = append(parts, p.nextToken().Value())
		}

		p.nextToken()

		return new(Expression(strings.Join(parts, ""))), nil
	}

	switch p.current().Type() {
	case tokens.TokenTypeString:
		return p.parseString()
	case tokens.TokenTypeNumber:
		return p.parseNumber()
	case tokens.TokenTypeTrue, tokens.TokenTypeFalse:
		return p.parseBoolean()
	case tokens.TokenTypeNull:
		return p.parseNull()
	default:
		token := p.current()
		p.nextToken()
		return new(Expression(token.Value())), nil
	}
}

func (p *Parser) parseAttribute() (*Attribute, error) {
	if p.current().Type() != tokens.TokenTypeIdentifier && p.current().Type() != tokens.TokenTypeString {
		return nil, p.wrapError(fmt.Errorf("unexpected token"))
	}

	token := p.current()
	p.nextToken()

	if _, err := p.expect(tokens.TokenTypeAssign); err != nil {
		return nil, err
	}

	value, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	return &Attribute{
		Name:  token.Value(),
		Value: value,
	}, nil
}

func (p *Parser) parseBlock() (*Block, error) {
	name, err := p.expect(tokens.TokenTypeIdentifier)
	if err != nil {
		return nil, err
	}

	var labels []string

	for p.current().Type() != tokens.TokenTypeLeftBrace {
		if p.current().Type() != tokens.TokenTypeString && p.current().Type() != tokens.TokenTypeIdentifier {
			return nil, p.wrapError(fmt.Errorf("unexpected token"))
		}

		label := p.current()
		p.nextToken()

		labels = append(labels, label.Value())
	}

	if _, err := p.expect(tokens.TokenTypeLeftBrace); err != nil {
		return nil, err
	}

	body, err := p.parseBody()
	if err != nil {
		return nil, err
	}

	return &Block{
		Name:   name.Value(),
		Labels: labels,
		Body:   *body,
	}, nil
}

func (p *Parser) parseBody() (*Body, error) {
	body := &Body{}

	for !p.match(tokens.TokenTypeRightBrace) {
		p.skipNewLines()

		if p.match(tokens.TokenTypeRightBrace) {
			break
		}

		if p.current().Type() == tokens.TokenTypeEOF {
			return nil, p.wrapError(fmt.Errorf("unexpected EOF"))
		}

		if p.peek(1).Type() == tokens.TokenTypeAssign {
			attr, err := p.parseAttribute()
			if err != nil {
				return nil, err
			}

			body.Nodes = append(body.Nodes, attr)
		} else {
			block, err := p.parseBlock()
			if err != nil {
				return nil, err
			}

			body.Nodes = append(body.Nodes, block)
		}

		p.skipNewLines()
	}

	return body, nil
}

func (p *Parser) parseFile() (*File, error) {
	body := &Body{}

	for p.current().Type() != tokens.TokenTypeEOF {
		p.skipNewLines()

		if p.current().Type() == tokens.TokenTypeEOF {
			break
		}

		var node Node

		if p.peek(1).Type() == tokens.TokenTypeAssign {
			attr, err := p.parseAttribute()
			if err != nil {
				return nil, err
			}
			node = attr
		} else {
			block, err := p.parseBlock()
			if err != nil {
				return nil, err
			}
			node = block
		}

		body.Nodes = append(body.Nodes, node)
		p.skipNewLines()
	}

	return &File{
		Body: *body,
	}, nil
}

func (p *Parser) Convert(file *File) (nodes.Object, error) {
	return convertBody(&file.Body)
}

func convertBody(body *Body) (nodes.Object, error) {
	result := make(nodes.Object)
	blockNames := make(map[string]bool)

	for _, node := range body.Nodes {
		switch n := node.(type) {
		case *Attribute:
			if _, exists := result[n.Name]; exists {
				return nil, fmt.Errorf("duplicate field name %q", n.Name)
			}

			value, err := convertValue(n.Value)
			if err != nil {
				return nil, fmt.Errorf("attribute %q: %w", n.Name, err)
			}

			result[n.Name] = value

		case *Block:
			value, err := convertBlock(n)
			if err != nil {
				return nil, fmt.Errorf("block %q: %w", n.Name, err)
			}

			if existing, ok := result[n.Name]; ok {
				if !blockNames[n.Name] {
					return nil, fmt.Errorf("duplicate field name %q", n.Name)
				}

				switch arr := existing.(type) {
				case nodes.Array:
					result[n.Name] = append(arr, value)

				default:
					result[n.Name] = nodes.Array{existing, value}
				}
			} else {
				result[n.Name] = value
				blockNames[n.Name] = true
			}

		default:
			return nil, nil
		}
	}

	return result, nil
}

func convertBlock(block *Block) (nodes.Value, error) {
	body, err := convertBody(&block.Body)
	if err != nil {
		return nil, err
	}

	var result nodes.Value = body

	for _, v := range slices.Backward(block.Labels) {
		result = nodes.Object{
			v: result,
		}
	}

	return result, nil
}

func convertValue(node Node) (nodes.Value, error) {
	switch n := node.(type) {
	case *String:
		return nodes.String(*n), nil

	case *Number:
		return nodes.Number(*n), nil

	case *Boolean:
		return nodes.Boolean(*n), nil

	case *Null:
		return nodes.Null{}, nil

	case *Expression:
		return nodes.Expression(*n), nil

	case *Array:
		result := make(nodes.Array, 0, len(*n))

		for _, element := range *n {
			value, err := convertValue(element)
			if err != nil {
				return nil, err
			}

			result = append(result, value)
		}

		return result, nil

	case *Object:
		result := make(nodes.Object)

		for _, attribute := range *n {
			if _, ok := attribute.(*Attribute); !ok {
				return nil, fmt.Errorf("unexpected node")
			}

			attr := attribute.(*Attribute)

			if _, exists := result[attr.Name]; exists {
				return nil, fmt.Errorf("duplicate object key %q", attr.Name)
			}

			value, err := convertValue(attr.Value)
			if err != nil {
				return nil, fmt.Errorf("object key %q: %w", attr.Name, err)
			}

			result[attr.Name] = value
		}

		return result, nil
	}

	return nil, nil
}
