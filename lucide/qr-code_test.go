package lucide

import (
	"os"
	"strings"
	"testing"

	x "github.com/bloxui/blox"
	"github.com/stretchr/testify/require"
)

func TestQrCode_SVGMatchesReference(t *testing.T) {
	// Render the component
	rendered := x.Render(QrCode())

	// Parse rendered SVG
	gotNode, err := parseXMLToNode(strings.NewReader(rendered))
	require.NoError(t, err, "parse rendered SVG")
	gotNode = normalizeNode(gotNode)

	// Load reference SVG
	f, err := os.Open("icons/qr-code.svg")
	require.NoError(t, err, "open reference SVG")
	defer f.Close()

	wantNode, err := parseXMLToNode(f)
	require.NoError(t, err, "parse reference SVG")
	wantNode = normalizeNode(wantNode)

	// Compare root element
	require.Equal(t, wantNode.Name.Local, gotNode.Name.Local, "root element name")
	require.Equal(t, wantNode.Attrs, gotNode.Attrs, "root attributes must match")

	// Compare recursively
	compareNodes(t, wantNode, gotNode)
}
