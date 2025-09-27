package lucide

import (
	html "github.com/plainkit/html"
)

// RectangleVertical creates a Rectangle Vertical Lucide icon.
func RectangleVertical(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-rectangle-vertical", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("12"), html.AHeight("20"), html.AX("6"), html.AY("2"), html.ARx("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
