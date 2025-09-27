package lucide

import (
	html "github.com/plainkit/html"
)

// BusFront creates a Bus Front Lucide icon.
func BusFront(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bus-front", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M4 6 2 7"))),
		html.Child(html.SvgPath(html.AD("M10 6h4"))),
		html.Child(html.SvgPath(html.AD("m22 7-2-1"))),
		html.Child(html.SvgRect(html.AWidth("16"), html.AHeight("16"), html.AX("4"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M4 11h16"))),
		html.Child(html.SvgPath(html.AD("M8 15h.01"))),
		html.Child(html.SvgPath(html.AD("M16 15h.01"))),
		html.Child(html.SvgPath(html.AD("M6 19v2"))),
		html.Child(html.SvgPath(html.AD("M18 21v-2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
