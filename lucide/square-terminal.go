package lucide

import (
	html "github.com/plainkit/html"
)

// SquareTerminal creates a Square Terminal Lucide icon.
func SquareTerminal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-terminal", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m7 11 2-2-2-2")),
		html.SvgPath(html.AD("M11 13h4")),
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"), html.ARy("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
