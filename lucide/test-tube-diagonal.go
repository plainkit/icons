package lucide

import (
	html "github.com/plainkit/html"
)

// TestTubeDiagonal creates a Test Tube Diagonal Lucide icon.
func TestTubeDiagonal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-test-tube-diagonal", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M21 7 6.82 21.18a2.83 2.83 0 0 1-3.99-.01a2.83 2.83 0 0 1 0-4L17 3")),
		html.SvgPath(html.AD("m16 2 6 6")),
		html.SvgPath(html.AD("M12 16H4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
