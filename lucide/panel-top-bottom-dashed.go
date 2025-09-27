package lucide

import (
	html "github.com/plainkit/html"
)

// PanelTopBottomDashed creates a Panel Top Bottom Dashed Lucide icon.
func PanelTopBottomDashed(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-panel-top-bottom-dashed", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M14 15h1"))),
		html.Child(html.SvgPath(html.AD("M14 9h1"))),
		html.Child(html.SvgPath(html.AD("M19 15h2"))),
		html.Child(html.SvgPath(html.AD("M19 9h2"))),
		html.Child(html.SvgPath(html.AD("M3 15h2"))),
		html.Child(html.SvgPath(html.AD("M3 9h2"))),
		html.Child(html.SvgPath(html.AD("M9 15h1"))),
		html.Child(html.SvgPath(html.AD("M9 9h1"))),
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
