package hcl

type Node interface {
	node()
}

type File struct {
	Body
}

type Body struct {
	Nodes []Node
}

type Block struct {
	Name   string
	Labels []string
	Body
}

type Attribute struct {
	Name  string
	Value Node
}

type String string
type Number string
type Boolean bool
type Null struct{}
type Expression string
type Array []Node
type Object []Node

func (s *String) node()     {}
func (n *Number) node()     {}
func (b *Boolean) node()    {}
func (n *Null) node()       {}
func (e *Expression) node() {}
func (a *Array) node()      {}
func (o *Object) node()     {}
func (o *Block) node()      {}
func (o *Attribute) node()  {}
func (o *File) node()       {}

func NewStringAttribute(name string, value string) *Attribute {
	return &Attribute{
		Name:  name,
		Value: new(String(value)),
	}
}

func NewNumberAttribute(name string, value string) *Attribute {
	return &Attribute{
		Name:  name,
		Value: new(Number(value)),
	}
}

func NewBooleanAttribute(name string, value bool) *Attribute {
	return &Attribute{
		Name:  name,
		Value: new(Boolean(value)),
	}
}

func NewNullAttribute(name string) *Attribute {
	return &Attribute{
		Name:  name,
		Value: new(Null),
	}
}

func NewExpressionAttribute(name string, value string) *Attribute {
	return &Attribute{
		Name:  name,
		Value: new(Expression(value)),
	}
}

func NewArrayAttribute(name string, value []Node) *Attribute {
	return &Attribute{
		Name:  name,
		Value: new(Array(value)),
	}
}

func NewObjectAttribute(name string, value []Node) *Attribute {
	return &Attribute{
		Name:  name,
		Value: new(Array(value)),
	}
}
