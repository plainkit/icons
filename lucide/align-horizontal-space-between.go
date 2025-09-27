package lucide

import (
	html "github.com/plainkit/html"
)

// AlignHorizontalSpaceBetween creates a Align Horizontal Space Between Lucide icon.
func AlignHorizontalSpaceBetween(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-align-horizontal-space-between", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("6"), html.AHeight("14"), html.AX("3"), html.AY("5"), html.ARx("2"))),
		html.Child(html.SvgRect(html.AWidth("6"), html.AHeight("10"), html.AX("15"), html.AY("7"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M3 2v20"))),
		html.Child(html.SvgPath(html.AD("M21 2v20"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
