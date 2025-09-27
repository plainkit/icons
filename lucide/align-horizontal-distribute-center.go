package lucide

import (
	html "github.com/plainkit/html"
)

// AlignHorizontalDistributeCenter creates a Align Horizontal Distribute Center Lucide icon.
func AlignHorizontalDistributeCenter(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-align-horizontal-distribute-center", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("6"), html.AHeight("14"), html.AX("4"), html.AY("5"), html.ARx("2"))),
		html.Child(html.SvgRect(html.AWidth("6"), html.AHeight("10"), html.AX("14"), html.AY("7"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M17 22v-5"))),
		html.Child(html.SvgPath(html.AD("M17 7V2"))),
		html.Child(html.SvgPath(html.AD("M7 22v-3"))),
		html.Child(html.SvgPath(html.AD("M7 5V2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
