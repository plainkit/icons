package lucide

import (
	html "github.com/plainkit/html"
)

// Underline creates a Underline Lucide icon.
func Underline(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-underline", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M6 4v6a6 6 0 0 0 12 0V4"))),
		html.Child(html.SvgLine(html.AX1("4"), html.AX2("20"), html.AY1("20"), html.AY2("20"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
