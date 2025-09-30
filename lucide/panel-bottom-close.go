package lucide

import (
	html "github.com/plainkit/html"
)

// PanelBottomClose creates a Panel Bottom Close Lucide icon.
func PanelBottomClose(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-panel-bottom-close", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M3 15h18")),
		html.SvgPath(html.AD("m15 8-3 3-3-3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
