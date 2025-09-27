package lucide

import (
	html "github.com/plainkit/html"
)

// AlignStartVertical creates a Align Start Vertical Lucide icon.
func AlignStartVertical(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-align-start-vertical", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("9"), html.AHeight("6"), html.AX("6"), html.AY("14"), html.ARx("2"))),
		html.Child(html.SvgRect(html.AWidth("16"), html.AHeight("6"), html.AX("6"), html.AY("4"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M2 2v20"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
