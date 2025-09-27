package lucide

import (
	html "github.com/plainkit/html"
)

// ShowerHead creates a Shower Head Lucide icon.
func ShowerHead(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-shower-head", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m4 4 2.5 2.5"))),
		html.Child(html.SvgPath(html.AD("M13.5 6.5a4.95 4.95 0 0 0-7 7"))),
		html.Child(html.SvgPath(html.AD("M15 5 5 15"))),
		html.Child(html.SvgPath(html.AD("M14 17v.01"))),
		html.Child(html.SvgPath(html.AD("M10 16v.01"))),
		html.Child(html.SvgPath(html.AD("M13 13v.01"))),
		html.Child(html.SvgPath(html.AD("M16 10v.01"))),
		html.Child(html.SvgPath(html.AD("M11 20v.01"))),
		html.Child(html.SvgPath(html.AD("M17 14v.01"))),
		html.Child(html.SvgPath(html.AD("M20 11v.01"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
