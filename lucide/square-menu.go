package lucide

import (
	html "github.com/plainkit/html"
)

// SquareMenu creates a Square Menu Lucide icon.
func SquareMenu(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-menu", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M7 8h10"))),
		html.Child(html.SvgPath(html.AD("M7 12h10"))),
		html.Child(html.SvgPath(html.AD("M7 16h10"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
