package lucide

import (
	html "github.com/plainkit/html"
)

// AlignVerticalDistributeStart creates a Align Vertical Distribute Start Lucide icon.
func AlignVerticalDistributeStart(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-align-vertical-distribute-start", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("14"), html.AHeight("6"), html.AX("5"), html.AY("14"), html.ARx("2")),
		html.SvgRect(html.AWidth("10"), html.AHeight("6"), html.AX("7"), html.AY("4"), html.ARx("2")),
		html.SvgPath(html.AD("M2 14h20")),
		html.SvgPath(html.AD("M2 4h20")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
