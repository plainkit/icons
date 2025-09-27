package lucide

import (
	html "github.com/plainkit/html"
)

// SquareCheck creates a Square Check Lucide icon.
func SquareCheck(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-check", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("m9 12 2 2 4-4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
