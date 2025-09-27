package lucide

import (
	html "github.com/plainkit/html"
)

// ChartNetwork creates a Chart Network Lucide icon.
func ChartNetwork(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chart-network", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m13.11 7.664 1.78 2.672"))),
		html.Child(html.SvgPath(html.AD("m14.162 12.788-3.324 1.424"))),
		html.Child(html.SvgPath(html.AD("m20 4-6.06 1.515"))),
		html.Child(html.SvgPath(html.AD("M3 3v16a2 2 0 0 0 2 2h16"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("6"), html.AR("2"))),
		html.Child(html.SvgCircle(html.ACx("16"), html.ACy("12"), html.AR("2"))),
		html.Child(html.SvgCircle(html.ACx("9"), html.ACy("15"), html.AR("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
