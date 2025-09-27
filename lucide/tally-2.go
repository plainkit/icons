package lucide

import (
	html "github.com/plainkit/html"
)

// Tally2 creates a Tally 2 Lucide icon.
func Tally2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-tally-2", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M4 4v16"))),
		html.Child(html.SvgPath(html.AD("M9 4v16"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
