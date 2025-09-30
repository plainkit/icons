package lucide

import (
	html "github.com/plainkit/html"
)

// AlignVerticalSpaceAround creates a Align Vertical Space Around Lucide icon.
func AlignVerticalSpaceAround(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-align-vertical-space-around", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("10"), html.AHeight("6"), html.AX("7"), html.AY("9"), html.ARx("2")),
		html.SvgPath(html.AD("M22 20H2")),
		html.SvgPath(html.AD("M22 4H2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
