package lucide

import (
	html "github.com/plainkit/html"
)

// SquareAsterisk creates a Square Asterisk Lucide icon.
func SquareAsterisk(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-asterisk", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M12 8v8"))),
		html.Child(html.SvgPath(html.AD("m8.5 14 7-4"))),
		html.Child(html.SvgPath(html.AD("m8.5 10 7 4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
