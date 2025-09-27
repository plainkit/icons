package lucide

import (
	html "github.com/plainkit/html"
)

// SquareX creates a Square X Lucide icon.
func SquareX(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-x", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"), html.ARy("2"))),
		html.Child(html.SvgPath(html.AD("m15 9-6 6"))),
		html.Child(html.SvgPath(html.AD("m9 9 6 6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
