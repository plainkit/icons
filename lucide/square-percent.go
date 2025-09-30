package lucide

import (
	html "github.com/plainkit/html"
)

// SquarePercent creates a Square Percent Lucide icon.
func SquarePercent(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-percent", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("m15 9-6 6")),
		html.SvgPath(html.AD("M9 9h.01")),
		html.SvgPath(html.AD("M15 15h.01")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
