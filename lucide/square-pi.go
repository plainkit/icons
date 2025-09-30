package lucide

import (
	html "github.com/plainkit/html"
)

// SquarePi creates a Square Pi Lucide icon.
func SquarePi(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-pi", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M7 7h10")),
		html.SvgPath(html.AD("M10 7v10")),
		html.SvgPath(html.AD("M16 17a2 2 0 0 1-2-2V7")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
