package lucide

import (
	html "github.com/plainkit/html"
)

// LayoutTemplate creates a Layout Template Lucide icon.
func LayoutTemplate(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-layout-template", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("7"), html.AX("3"), html.AY("3"), html.ARx("1")),
		html.SvgRect(html.AWidth("9"), html.AHeight("7"), html.AX("3"), html.AY("14"), html.ARx("1")),
		html.SvgRect(html.AWidth("5"), html.AHeight("7"), html.AX("16"), html.AY("14"), html.ARx("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
