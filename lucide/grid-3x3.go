package lucide

import (
	html "github.com/plainkit/html"
)

// Grid3x3 creates a Grid 3x3 Lucide icon.
func Grid3x3(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-grid-3x3", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M3 9h18")),
		html.SvgPath(html.AD("M3 15h18")),
		html.SvgPath(html.AD("M9 3v18")),
		html.SvgPath(html.AD("M15 3v18")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
