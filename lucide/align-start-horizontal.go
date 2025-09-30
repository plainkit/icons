package lucide

import (
	html "github.com/plainkit/html"
)

// AlignStartHorizontal creates a Align Start Horizontal Lucide icon.
func AlignStartHorizontal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-align-start-horizontal", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("6"), html.AHeight("16"), html.AX("4"), html.AY("6"), html.ARx("2")),
		html.SvgRect(html.AWidth("6"), html.AHeight("9"), html.AX("14"), html.AY("6"), html.ARx("2")),
		html.SvgPath(html.AD("M22 2H2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
