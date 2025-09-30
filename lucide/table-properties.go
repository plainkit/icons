package lucide

import (
	html "github.com/plainkit/html"
)

// TableProperties creates a Table Properties Lucide icon.
func TableProperties(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-table-properties", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15 3v18")),
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M21 9H3")),
		html.SvgPath(html.AD("M21 15H3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
