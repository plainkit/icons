package lucide

import (
	html "github.com/plainkit/html"
)

// AlignVerticalSpaceBetween creates a Align Vertical Space Between Lucide icon.
func AlignVerticalSpaceBetween(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-align-vertical-space-between", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("14"), html.AHeight("6"), html.AX("5"), html.AY("15"), html.ARx("2"))),
		html.Child(html.SvgRect(html.AWidth("10"), html.AHeight("6"), html.AX("7"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M2 21h20"))),
		html.Child(html.SvgPath(html.AD("M2 3h20"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
