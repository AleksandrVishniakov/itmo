package xml

import (
	"fmt"
	"io"
	"parser/nodes"
	"strings"
)

const header = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>"

type Writer struct {
	w io.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w}
}

func (w *Writer) Write(node nodes.Value) error {
	if _, err := io.WriteString(w.w, header+"\n"); err != nil {
		return err
	}

	return w.writeAST("config", node, 0)
}

func (w *Writer) writeAST(name string, node nodes.Value, depth int) error {
	indent := strings.Repeat(" ", depth*2)
	name = strings.ReplaceAll(name, " ", "-")

	switch n := node.(type) {
	case nodes.Object:
		if _, err := fmt.Fprintf(w.w, "%s<%s>\n", indent, name); err != nil {
			return err
		}

		for key, value := range n {
			if err := w.writeAST(key, value, depth+1); err != nil {
				return err
			}
		}

		_, err := fmt.Fprintf(w.w, "%s</%s>\n", indent, name)
		return err

	case nodes.Array:
		if _, err := fmt.Fprintf(w.w, "%s<%s>\n", indent, name); err != nil {
			return err
		}

		for _, value := range n {
			if err := w.writeAST("item", value, depth+1); err != nil {
				return err
			}
		}

		_, err := fmt.Fprintf(w.w, "%s</%s>\n", indent, name)
		return err

	case nodes.String:
		return w.writeElement(name, string(n), depth)

	case nodes.Number:
		return w.writeElement(name, string(n), depth)

	case nodes.Boolean:
		return w.writeElement(name, fmt.Sprint(bool(n)), depth)

	case nodes.Null:
		_, err := fmt.Fprintf(w.w, "%s<%s/>\n", indent, name)
		return err

	case nodes.Expression:
		return w.writeElement(name, string(n), depth)
	}

	return nil
}

func (w *Writer) writeElement(name, value string, depth int) error {
	indent := strings.Repeat(" ", depth*2)

	_, err := fmt.Fprintf(w.w, "%s<%s>%s</%s>\n", indent, name, value, name)

	return err
}
