package lucide

import (
	html "github.com/plainkit/html"
)

// Dices creates a Dices Lucide icon.
func Dices(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-dices", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("12"), html.AHeight("12"), html.AX("2"), html.AY("10"), html.ARx("2"), html.ARy("2")),
		html.SvgPath(html.AD("m17.92 14 3.5-3.5a2.24 2.24 0 0 0 0-3l-5-4.92a2.24 2.24 0 0 0-3 0L10 6")),
		html.SvgPath(html.AD("M6 18h.01")),
		html.SvgPath(html.AD("M10 14h.01")),
		html.SvgPath(html.AD("M15 6h.01")),
		html.SvgPath(html.AD("M18 9h.01")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
