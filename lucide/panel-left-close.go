package lucide

import (
	html "github.com/plainkit/html"
)

// PanelLeftClose creates a Panel Left Close Lucide icon.
func PanelLeftClose(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-panel-left-close", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M9 3v18"))),
		html.Child(html.SvgPath(html.AD("m16 15-3-3 3-3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
