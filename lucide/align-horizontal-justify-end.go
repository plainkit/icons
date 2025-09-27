package lucide

import (
	html "github.com/plainkit/html"
)

// AlignHorizontalJustifyEnd creates a Align Horizontal Justify End Lucide icon.
func AlignHorizontalJustifyEnd(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-align-horizontal-justify-end", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("6"), html.AHeight("14"), html.AX("2"), html.AY("5"), html.ARx("2"))),
		html.Child(html.SvgRect(html.AWidth("6"), html.AHeight("10"), html.AX("12"), html.AY("7"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M22 2v20"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
