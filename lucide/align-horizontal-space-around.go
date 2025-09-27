package lucide

import (
	html "github.com/plainkit/html"
)

// AlignHorizontalSpaceAround creates a Align Horizontal Space Around Lucide icon.
func AlignHorizontalSpaceAround(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-align-horizontal-space-around", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("6"), html.AHeight("10"), html.AX("9"), html.AY("7"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M4 22V2"))),
		html.Child(html.SvgPath(html.AD("M20 22V2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
