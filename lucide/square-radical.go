package lucide

import (
	html "github.com/plainkit/html"
)

// SquareRadical creates a Square Radical Lucide icon.
func SquareRadical(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-radical", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M7 12h2l2 5 2-10h4")),
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
