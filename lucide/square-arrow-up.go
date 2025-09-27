package lucide

import (
	html "github.com/plainkit/html"
)

// SquareArrowUp creates a Square Arrow Up Lucide icon.
func SquareArrowUp(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-arrow-up", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("m16 12-4-4-4 4"))),
		html.Child(html.SvgPath(html.AD("M12 16V8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
