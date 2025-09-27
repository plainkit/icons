package lucide

import (
	html "github.com/plainkit/html"
)

// Router creates a Router Lucide icon.
func Router(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-router", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("20"), html.AHeight("8"), html.AX("2"), html.AY("14"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M6.01 18H6"))),
		html.Child(html.SvgPath(html.AD("M10.01 18H10"))),
		html.Child(html.SvgPath(html.AD("M15 10v4"))),
		html.Child(html.SvgPath(html.AD("M17.84 7.17a4 4 0 0 0-5.66 0"))),
		html.Child(html.SvgPath(html.AD("M20.66 4.34a8 8 0 0 0-11.31 0"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
