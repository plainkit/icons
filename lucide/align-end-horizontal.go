package lucide

import (
	html "github.com/plainkit/html"
)

// AlignEndHorizontal creates a Align End Horizontal Lucide icon.
func AlignEndHorizontal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-align-end-horizontal", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("6"), html.AHeight("16"), html.AX("4"), html.AY("2"), html.ARx("2")),
		html.SvgRect(html.AWidth("6"), html.AHeight("9"), html.AX("14"), html.AY("9"), html.ARx("2")),
		html.SvgPath(html.AD("M22 22H2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
