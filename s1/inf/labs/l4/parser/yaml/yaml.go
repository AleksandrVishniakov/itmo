package yaml

import (
	"io"
	"parser/nodes"
)

type YAML struct{}

func (y YAML) FromAST(w io.Writer, root nodes.Value) error {
	writer := NewWriter(w)
	return writer.Write(root)
}
