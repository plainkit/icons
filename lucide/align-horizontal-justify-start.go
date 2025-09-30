package lucide

import (
	html "github.com/plainkit/html"
)

// AlignHorizontalJustifyStart creates a Align Horizontal Justify Start Lucide icon.
func AlignHorizontalJustifyStart(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-align-horizontal-justify-start", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("6"), html.AHeight("14"), html.AX("6"), html.AY("5"), html.ARx("2")),
		html.SvgRect(html.AWidth("6"), html.AHeight("10"), html.AX("16"), html.AY("7"), html.ARx("2")),
		html.SvgPath(html.AD("M2 2v20")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
