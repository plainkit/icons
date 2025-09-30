package lucide

import (
	html "github.com/plainkit/html"
)

// PanelRightDashed creates a Panel Right Dashed Lucide icon.
func PanelRightDashed(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-panel-right-dashed", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M15 14v1")),
		html.SvgPath(html.AD("M15 19v2")),
		html.SvgPath(html.AD("M15 3v2")),
		html.SvgPath(html.AD("M15 9v1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
