package lucide

import (
	"encoding/xml"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// xmlNode is a minimal DOM-like node for structural comparison
type xmlNode struct {
	Name     xml.Name
	Attrs    map[string]string
	Children []xmlNode
	Text     string
}

func parseXMLToNode(r io.Reader) (xmlNode, error) {
	dec := xml.NewDecoder(r)
	// read until first StartElement
	for {
		tok, err := dec.Token()
		if err != nil {
			return xmlNode{}, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			return readElement(dec, t), nil
		}
	}
}

func readElement(dec *xml.Decoder, start xml.StartElement) xmlNode {
	n := xmlNode{
		Name:  start.Name,
		Attrs: make(map[string]string, len(start.Attr)),
	}
	for _, a := range start.Attr {
		// Normalize attribute names to their local part
		key := a.Name.Local
		n.Attrs[key] = strings.TrimSpace(a.Value)
	}

	for {
		tok, err := dec.Token()
		if err != nil {
			// unexpected EOF is treated as termination
			break
		}
		switch t := tok.(type) {
		case xml.StartElement:
			child := readElement(dec, t)
			n.Children = append(n.Children, child)
		case xml.EndElement:
			if t.Name.Local == start.Name.Local {
				return n
			}
		case xml.CharData:
			txt := strings.TrimSpace(string([]byte(t)))
			if txt != "" {
				if n.Text != "" {
					n.Text += " "
				}
				n.Text += txt
			}
		}
	}
	return n
}

func normalizeNode(n xmlNode) xmlNode {
	// Remove attributes that are intentionally added by buildLucideArgs
	// but not present in the reference SVGs (class on root)
	// We only drop class when present on the root <svg>.
	if n.Name.Local == "svg" {
		delete(n.Attrs, "class")
	}
	// No reordering of children: order matters
	// Attribute order is irrelevant due to map representation
	for i := range n.Children {
		n.Children[i] = normalizeNode(n.Children[i])
	}
	return n
}

func compareNodes(t *testing.T, want, got xmlNode) {
	t.Helper()
	require.Equal(t, want.Name.Local, got.Name.Local, "element name")
	require.Equal(t, want.Attrs, got.Attrs, "attributes for <%s>", want.Name.Local)
	require.Equal(t, want.Text, got.Text, "text for <%s>", want.Name.Local)
	require.Equal(t, len(want.Children), len(got.Children), "children count for <%s>", want.Name.Local)
	for i := range want.Children {
		compareNodes(t, want.Children[i], got.Children[i])
	}
}
