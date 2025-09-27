package lucide

import (
	html "github.com/plainkit/html"
)

// SquareActivity creates a Square Activity Lucide icon.
func SquareActivity(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-activity", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M17 12h-2l-2 5-2-10-2 5H7"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
