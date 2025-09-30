package lucide

import (
	html "github.com/plainkit/html"
)

// Grid2x2 creates a Grid 2x2 Lucide icon.
func Grid2x2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-grid-2x2", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 3v18")),
		html.SvgPath(html.AD("M3 12h18")),
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
