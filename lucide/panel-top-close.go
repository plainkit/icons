package lucide

import (
	html "github.com/plainkit/html"
)

// PanelTopClose creates a Panel Top Close Lucide icon.
func PanelTopClose(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-panel-top-close", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M3 9h18")),
		html.SvgPath(html.AD("m9 16 3-3 3 3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
