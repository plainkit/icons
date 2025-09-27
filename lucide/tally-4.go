package lucide

import (
	html "github.com/plainkit/html"
)

// Tally4 creates a Tally 4 Lucide icon.
func Tally4(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-tally-4", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M4 4v16"))),
		html.Child(html.SvgPath(html.AD("M9 4v16"))),
		html.Child(html.SvgPath(html.AD("M14 4v16"))),
		html.Child(html.SvgPath(html.AD("M19 4v16"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
