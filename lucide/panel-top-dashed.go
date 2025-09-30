package lucide

import (
	html "github.com/plainkit/html"
)

// PanelTopDashed creates a Panel Top Dashed Lucide icon.
func PanelTopDashed(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-panel-top-dashed", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M14 9h1")),
		html.SvgPath(html.AD("M19 9h2")),
		html.SvgPath(html.AD("M3 9h2")),
		html.SvgPath(html.AD("M9 9h1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
