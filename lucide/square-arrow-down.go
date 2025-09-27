package lucide

import (
	html "github.com/plainkit/html"
)

// SquareArrowDown creates a Square Arrow Down Lucide icon.
func SquareArrowDown(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-arrow-down", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M12 8v8"))),
		html.Child(html.SvgPath(html.AD("m8 12 4 4 4-4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
