package lucide

import (
	html "github.com/plainkit/html"
)

// SquarePlus creates a Square Plus Lucide icon.
func SquarePlus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-plus", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M8 12h8"))),
		html.Child(html.SvgPath(html.AD("M12 8v8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
