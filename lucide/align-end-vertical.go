package lucide

import (
	html "github.com/plainkit/html"
)

// AlignEndVertical creates a Align End Vertical Lucide icon.
func AlignEndVertical(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-align-end-vertical", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("16"), html.AHeight("6"), html.AX("2"), html.AY("4"), html.ARx("2"))),
		html.Child(html.SvgRect(html.AWidth("9"), html.AHeight("6"), html.AX("9"), html.AY("14"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M22 22V2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
