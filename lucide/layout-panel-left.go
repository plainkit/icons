package lucide

import (
	html "github.com/plainkit/html"
)

// LayoutPanelLeft creates a Layout Panel Left Lucide icon.
func LayoutPanelLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-layout-panel-left", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("7"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("1"))),
		html.Child(html.SvgRect(html.AWidth("7"), html.AHeight("7"), html.AX("14"), html.AY("3"), html.ARx("1"))),
		html.Child(html.SvgRect(html.AWidth("7"), html.AHeight("7"), html.AX("14"), html.AY("14"), html.ARx("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
