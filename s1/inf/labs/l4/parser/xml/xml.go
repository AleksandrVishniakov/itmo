package xml

import (
	"io"
	"parser/nodes"
)

type XML struct{}

func (x XML) FromAST(w io.Writer, root nodes.Value) error {
	writer := NewWriter(w)
	return writer.Write(root)
}
