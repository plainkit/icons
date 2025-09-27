package lucide

import (
	html "github.com/plainkit/html"
)

// PanelTopOpen creates a Panel Top Open Lucide icon.
func PanelTopOpen(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-panel-top-open", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M3 9h18"))),
		html.Child(html.SvgPath(html.AD("m15 14-3 3-3-3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
