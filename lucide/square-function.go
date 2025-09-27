package lucide

import (
	html "github.com/plainkit/html"
)

// SquareFunction creates a Square Function Lucide icon.
func SquareFunction(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-function", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"), html.ARy("2"))),
		html.Child(html.SvgPath(html.AD("M9 17c2 0 2.8-1 2.8-2.8V10c0-2 1-3.3 3.2-3"))),
		html.Child(html.SvgPath(html.AD("M9 11.2h5.7"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
