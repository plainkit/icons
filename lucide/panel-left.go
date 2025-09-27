package lucide

import (
	html "github.com/plainkit/html"
)

// PanelLeft creates a Panel Left Lucide icon.
func PanelLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-panel-left", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M9 3v18"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
