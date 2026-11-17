package nodes

type Value interface {
	valueNode()
}

type Object map[string]Value

type Array []Value

type String string
type Number string
type Expression string
type Boolean bool
type Null struct{}

func (Object) valueNode()     {}
func (Array) valueNode()      {}
func (String) valueNode()     {}
func (Number) valueNode()     {}
func (Expression) valueNode() {}
func (Boolean) valueNode()    {}
func (Null) valueNode()       {}
