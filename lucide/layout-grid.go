package lucide

import (
	html "github.com/plainkit/html"
)

// LayoutGrid creates a Layout Grid Lucide icon.
func LayoutGrid(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-layout-grid", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("7"), html.AHeight("7"), html.AX("3"), html.AY("3"), html.ARx("1"))),
		html.Child(html.SvgRect(html.AWidth("7"), html.AHeight("7"), html.AX("14"), html.AY("3"), html.ARx("1"))),
		html.Child(html.SvgRect(html.AWidth("7"), html.AHeight("7"), html.AX("14"), html.AY("14"), html.ARx("1"))),
		html.Child(html.SvgRect(html.AWidth("7"), html.AHeight("7"), html.AX("3"), html.AY("14"), html.ARx("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
