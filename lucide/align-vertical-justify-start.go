package lucide

import (
	html "github.com/plainkit/html"
)

// AlignVerticalJustifyStart creates a Align Vertical Justify Start Lucide icon.
func AlignVerticalJustifyStart(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-align-vertical-justify-start", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("14"), html.AHeight("6"), html.AX("5"), html.AY("16"), html.ARx("2"))),
		html.Child(html.SvgRect(html.AWidth("10"), html.AHeight("6"), html.AX("7"), html.AY("6"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M2 2h20"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
