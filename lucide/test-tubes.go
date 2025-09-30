package lucide

import (
	html "github.com/plainkit/html"
)

// TestTubes creates a Test Tubes Lucide icon.
func TestTubes(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-test-tubes", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M9 2v17.5A2.5 2.5 0 0 1 6.5 22A2.5 2.5 0 0 1 4 19.5V2")),
		html.SvgPath(html.AD("M20 2v17.5a2.5 2.5 0 0 1-2.5 2.5a2.5 2.5 0 0 1-2.5-2.5V2")),
		html.SvgPath(html.AD("M3 2h7")),
		html.SvgPath(html.AD("M14 2h7")),
		html.SvgPath(html.AD("M9 16H4")),
		html.SvgPath(html.AD("M20 16h-5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
