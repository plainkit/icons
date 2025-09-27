package lucide

import (
	html "github.com/plainkit/html"
)

// AlignVerticalJustifyCenter creates a Align Vertical Justify Center Lucide icon.
func AlignVerticalJustifyCenter(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-align-vertical-justify-center", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("14"), html.AHeight("6"), html.AX("5"), html.AY("16"), html.ARx("2"))),
		html.Child(html.SvgRect(html.AWidth("10"), html.AHeight("6"), html.AX("7"), html.AY("2"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M2 12h20"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
