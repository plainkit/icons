package lucide

import (
	html "github.com/plainkit/html"
)

// VectorSquare creates a Vector Square Lucide icon.
func VectorSquare(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-vector-square", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M19.5 7a24 24 0 0 1 0 10"))),
		html.Child(html.SvgPath(html.AD("M4.5 7a24 24 0 0 0 0 10"))),
		html.Child(html.SvgPath(html.AD("M7 19.5a24 24 0 0 0 10 0"))),
		html.Child(html.SvgPath(html.AD("M7 4.5a24 24 0 0 1 10 0"))),
		html.Child(html.SvgRect(html.AWidth("5"), html.AHeight("5"), html.AX("17"), html.AY("17"), html.ARx("1"))),
		html.Child(html.SvgRect(html.AWidth("5"), html.AHeight("5"), html.AX("17"), html.AY("2"), html.ARx("1"))),
		html.Child(html.SvgRect(html.AWidth("5"), html.AHeight("5"), html.AX("2"), html.AY("17"), html.ARx("1"))),
		html.Child(html.SvgRect(html.AWidth("5"), html.AHeight("5"), html.AX("2"), html.AY("2"), html.ARx("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
