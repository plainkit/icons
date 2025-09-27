package lucide

import (
	html "github.com/plainkit/html"
)

// SquarePilcrow creates a Square Pilcrow Lucide icon.
func SquarePilcrow(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-pilcrow", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M12 12H9.5a2.5 2.5 0 0 1 0-5H17"))),
		html.Child(html.SvgPath(html.AD("M12 7v10"))),
		html.Child(html.SvgPath(html.AD("M16 7v10"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
