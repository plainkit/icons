package lucide

import (
	html "github.com/plainkit/html"
)

// PanelTop creates a Panel Top Lucide icon.
func PanelTop(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-panel-top", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M3 9h18")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
