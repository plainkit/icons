package lucide

import (
	html "github.com/plainkit/html"
)

// AlignVerticalDistributeCenter creates a Align Vertical Distribute Center Lucide icon.
func AlignVerticalDistributeCenter(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-align-vertical-distribute-center", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M22 17h-3")),
		html.SvgPath(html.AD("M22 7h-5")),
		html.SvgPath(html.AD("M5 17H2")),
		html.SvgPath(html.AD("M7 7H2")),
		html.SvgRect(html.AWidth("14"), html.AHeight("6"), html.AX("5"), html.AY("14"), html.ARx("2")),
		html.SvgRect(html.AWidth("10"), html.AHeight("6"), html.AX("7"), html.AY("4"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
