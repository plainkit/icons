package lucide

import (
	html "github.com/plainkit/html"
)

// Italic creates a Italic Lucide icon.
func Italic(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-italic", args)
	children := []html.SvgArg{
		html.Child(html.SvgLine(html.AX1("19"), html.AX2("10"), html.AY1("4"), html.AY2("4"))),
		html.Child(html.SvgLine(html.AX1("14"), html.AX2("5"), html.AY1("20"), html.AY2("20"))),
		html.Child(html.SvgLine(html.AX1("15"), html.AX2("9"), html.AY1("4"), html.AY2("20"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
