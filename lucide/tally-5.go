package lucide

import (
	html "github.com/plainkit/html"
)

// Tally5 creates a Tally 5 Lucide icon.
func Tally5(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-tally-5", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M4 4v16"))),
		html.Child(html.SvgPath(html.AD("M9 4v16"))),
		html.Child(html.SvgPath(html.AD("M14 4v16"))),
		html.Child(html.SvgPath(html.AD("M19 4v16"))),
		html.Child(html.SvgPath(html.AD("M22 6 2 18"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
