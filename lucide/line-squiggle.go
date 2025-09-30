package lucide

import (
	html "github.com/plainkit/html"
)

// LineSquiggle creates a Line Squiggle Lucide icon.
func LineSquiggle(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-line-squiggle", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M7 3.5c5-2 7 2.5 3 4C1.5 10 2 15 5 16c5 2 9-10 14-7s.5 13.5-4 12c-5-2.5.5-11 6-2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
