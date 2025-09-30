package lucide

import (
	html "github.com/plainkit/html"
)

// Table creates a Table Lucide icon.
func Table(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-table", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 3v18")),
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M3 9h18")),
		html.SvgPath(html.AD("M3 15h18")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
