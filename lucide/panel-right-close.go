package lucide

import (
	html "github.com/plainkit/html"
)

// PanelRightClose creates a Panel Right Close Lucide icon.
func PanelRightClose(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-panel-right-close", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M15 3v18")),
		html.SvgPath(html.AD("m8 9 3 3-3 3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
