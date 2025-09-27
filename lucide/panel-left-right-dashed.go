package lucide

import (
	html "github.com/plainkit/html"
)

// PanelLeftRightDashed creates a Panel Left Right Dashed Lucide icon.
func PanelLeftRightDashed(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-panel-left-right-dashed", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M15 10V9"))),
		html.Child(html.SvgPath(html.AD("M15 15v-1"))),
		html.Child(html.SvgPath(html.AD("M15 21v-2"))),
		html.Child(html.SvgPath(html.AD("M15 5V3"))),
		html.Child(html.SvgPath(html.AD("M9 10V9"))),
		html.Child(html.SvgPath(html.AD("M9 15v-1"))),
		html.Child(html.SvgPath(html.AD("M9 21v-2"))),
		html.Child(html.SvgPath(html.AD("M9 5V3"))),
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
