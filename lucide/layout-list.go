package lucide

import (
	html "github.com/plainkit/html"
)

// LayoutList creates a Layout List Lucide icon.
func LayoutList(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-layout-list", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("7"), html.AHeight("7"), html.AX("3"), html.AY("3"), html.ARx("1")),
		html.SvgRect(html.AWidth("7"), html.AHeight("7"), html.AX("3"), html.AY("14"), html.ARx("1")),
		html.SvgPath(html.AD("M14 4h7")),
		html.SvgPath(html.AD("M14 9h7")),
		html.SvgPath(html.AD("M14 15h7")),
		html.SvgPath(html.AD("M14 20h7")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
