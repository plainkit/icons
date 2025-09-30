package lucide

import (
	html "github.com/plainkit/html"
)

// PanelRight creates a Panel Right Lucide icon.
func PanelRight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-panel-right", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M15 3v18")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
