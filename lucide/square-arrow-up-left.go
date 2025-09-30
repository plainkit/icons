package lucide

import (
	html "github.com/plainkit/html"
)

// SquareArrowUpLeft creates a Square Arrow Up Left Lucide icon.
func SquareArrowUpLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-arrow-up-left", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M8 16V8h8")),
		html.SvgPath(html.AD("M16 16 8 8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
