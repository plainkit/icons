package lucide

import (
	html "github.com/plainkit/html"
)

// TestTube creates a Test Tube Lucide icon.
func TestTube(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-test-tube", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M14.5 2v17.5c0 1.4-1.1 2.5-2.5 2.5c-1.4 0-2.5-1.1-2.5-2.5V2")),
		html.SvgPath(html.AD("M8.5 2h7")),
		html.SvgPath(html.AD("M14.5 16h-5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
