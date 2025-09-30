package lucide

import (
	html "github.com/plainkit/html"
)

// PanelLeftOpen creates a Panel Left Open Lucide icon.
func PanelLeftOpen(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-panel-left-open", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M9 3v18")),
		html.SvgPath(html.AD("m14 9 3 3-3 3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
