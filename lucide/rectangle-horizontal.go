package lucide

import (
	html "github.com/plainkit/html"
)

// RectangleHorizontal creates a Rectangle Horizontal Lucide icon.
func RectangleHorizontal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-rectangle-horizontal", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("20"), html.AHeight("12"), html.AX("2"), html.AY("6"), html.ARx("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
