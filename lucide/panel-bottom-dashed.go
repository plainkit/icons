package lucide

import (
	html "github.com/plainkit/html"
)

// PanelBottomDashed creates a Panel Bottom Dashed Lucide icon.
func PanelBottomDashed(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-panel-bottom-dashed", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M14 15h1"))),
		html.Child(html.SvgPath(html.AD("M19 15h2"))),
		html.Child(html.SvgPath(html.AD("M3 15h2"))),
		html.Child(html.SvgPath(html.AD("M9 15h1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
