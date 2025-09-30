package lucide

import (
	html "github.com/plainkit/html"
)

// LayoutPanelTop creates a Layout Panel Top Lucide icon.
func LayoutPanelTop(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-layout-panel-top", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("7"), html.AX("3"), html.AY("3"), html.ARx("1")),
		html.SvgRect(html.AWidth("7"), html.AHeight("7"), html.AX("3"), html.AY("14"), html.ARx("1")),
		html.SvgRect(html.AWidth("7"), html.AHeight("7"), html.AX("14"), html.AY("14"), html.ARx("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
