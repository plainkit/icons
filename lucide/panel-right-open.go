package lucide

import (
	html "github.com/plainkit/html"
)

// PanelRightOpen creates a Panel Right Open Lucide icon.
func PanelRightOpen(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-panel-right-open", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M15 3v18")),
		html.SvgPath(html.AD("m10 15-3-3 3-3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
