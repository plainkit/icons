package lucide

import (
	html "github.com/plainkit/html"
)

// PanelBottomOpen creates a Panel Bottom Open Lucide icon.
func PanelBottomOpen(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-panel-bottom-open", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M3 15h18"))),
		html.Child(html.SvgPath(html.AD("m9 10 3-3 3 3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
