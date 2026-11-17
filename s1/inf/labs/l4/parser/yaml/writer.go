package yaml

import (
	"fmt"
	"io"
	"parser/nodes"
	"strconv"
	"strings"
)

type Writer struct {
	w io.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{
		w: w,
	}
}

func (w *Writer) Write(node nodes.Value) error {
	return w.writeAST(node, 0)
}

func (w *Writer) writeAST(node nodes.Value, depth int) error {
	indent := strings.Repeat(" ", depth*2)

	switch n := node.(type) {
	case nodes.Object:
		if len(n) == 0 {
			_, err := io.WriteString(w.w, " {}")
			return err
		}

		first := true
		for key, value := range n {
			if !first {
				if _, err := io.WriteString(w.w, "\n"); err != nil {
					return err
				}
			}
			first = false

			if _, err := fmt.Fprintf(w.w, "%s%s:", indent, key); err != nil {
				return err
			}

			switch v := value.(type) {
			case nodes.Array:
				if len(v) > 0 {
					if _, err := io.WriteString(w.w, "\n"); err != nil {
						return err
					}
				}
			case nodes.Object:
				if len(v) > 0 {
					if _, err := io.WriteString(w.w, "\n"); err != nil {
						return err
					}
				}
			default:
				if _, err := io.WriteString(w.w, " "); err != nil {
					return err
				}
			}

			if err := w.writeAST(value, depth+1); err != nil {
				return err
			}
		}

	case nodes.Array:
		if len(n) == 0 {
			_, err := io.WriteString(w.w, " []")
			return err
		}

		for i, value := range n {
			if i > 0 {
				if _, err := io.WriteString(w.w, "\n"); err != nil {
					return err
				}
			}

			if _, err := fmt.Fprintf(w.w, "%s- ", indent); err != nil {
				return err
			}

			switch value.(type) {
			case nodes.Object, nodes.Array:
				if _, err := io.WriteString(w.w, "\n"); err != nil {
					return err
				}
			}

			if err := w.writeAST(value, depth+1); err != nil {
				return err
			}
		}

	case nodes.String:
		_, err := io.WriteString(w.w, strconv.Quote(string(n)))
		return err

	case nodes.Number:
		_, err := io.WriteString(w.w, string(n))
		return err

	case nodes.Boolean:
		_, err := fmt.Fprint(w.w, bool(n))
		return err

	case nodes.Null:
		_, err := io.WriteString(w.w, "null")
		return err

	case nodes.Expression:
		_, err := io.WriteString(w.w, "${"+string(n)+"}")
		return err
	}

	return nil
}
