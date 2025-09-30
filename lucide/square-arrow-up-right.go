package lucide

import (
	html "github.com/plainkit/html"
)

// SquareArrowUpRight creates a Square Arrow Up Right Lucide icon.
func SquareArrowUpRight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-arrow-up-right", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M8 8h8v8")),
		html.SvgPath(html.AD("m8 16 8-8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
