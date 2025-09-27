package lucide

import (
	html "github.com/plainkit/html"
)

// RectangleEllipsis creates a Rectangle Ellipsis Lucide icon.
func RectangleEllipsis(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-rectangle-ellipsis", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("20"), html.AHeight("12"), html.AX("2"), html.AY("6"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M12 12h.01"))),
		html.Child(html.SvgPath(html.AD("M17 12h.01"))),
		html.Child(html.SvgPath(html.AD("M7 12h.01"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
