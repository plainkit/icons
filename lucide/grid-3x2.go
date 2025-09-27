package lucide

import (
	html "github.com/plainkit/html"
)

// Grid3x2 creates a Grid 3x2 Lucide icon.
func Grid3x2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-grid-3x2", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M15 3v18"))),
		html.Child(html.SvgPath(html.AD("M3 12h18"))),
		html.Child(html.SvgPath(html.AD("M9 3v18"))),
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
