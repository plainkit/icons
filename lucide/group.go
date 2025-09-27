package lucide

import (
	html "github.com/plainkit/html"
)

// Group creates a Group Lucide icon.
func Group(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-group", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3 7V5c0-1.1.9-2 2-2h2"))),
		html.Child(html.SvgPath(html.AD("M17 3h2c1.1 0 2 .9 2 2v2"))),
		html.Child(html.SvgPath(html.AD("M21 17v2c0 1.1-.9 2-2 2h-2"))),
		html.Child(html.SvgPath(html.AD("M7 21H5c-1.1 0-2-.9-2-2v-2"))),
		html.Child(html.SvgRect(html.AWidth("7"), html.AHeight("5"), html.AX("7"), html.AY("7"), html.ARx("1"))),
		html.Child(html.SvgRect(html.AWidth("7"), html.AHeight("5"), html.AX("10"), html.AY("12"), html.ARx("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
