package lucide

import (
	html "github.com/plainkit/html"
)

// PanelLeftDashed creates a Panel Left Dashed Lucide icon.
func PanelLeftDashed(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-panel-left-dashed", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M9 14v1"))),
		html.Child(html.SvgPath(html.AD("M9 19v2"))),
		html.Child(html.SvgPath(html.AD("M9 3v2"))),
		html.Child(html.SvgPath(html.AD("M9 9v1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
