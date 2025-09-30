package lucide

import (
	html "github.com/plainkit/html"
)

// LayoutDashboard creates a Layout Dashboard Lucide icon.
func LayoutDashboard(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-layout-dashboard", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("7"), html.AHeight("9"), html.AX("3"), html.AY("3"), html.ARx("1")),
		html.SvgRect(html.AWidth("7"), html.AHeight("5"), html.AX("14"), html.AY("3"), html.ARx("1")),
		html.SvgRect(html.AWidth("7"), html.AHeight("9"), html.AX("14"), html.AY("12"), html.ARx("1")),
		html.SvgRect(html.AWidth("7"), html.AHeight("5"), html.AX("3"), html.AY("16"), html.ARx("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
