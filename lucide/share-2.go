package lucide

import (
	html "github.com/plainkit/html"
)

// Share2 creates a Share 2 Lucide icon.
func Share2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-share-2", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("18"), html.ACy("5"), html.AR("3")),
		html.SvgCircle(html.ACx("6"), html.ACy("12"), html.AR("3")),
		html.SvgCircle(html.ACx("18"), html.ACy("19"), html.AR("3")),
		html.SvgLine(html.AX1("8.59"), html.AX2("15.42"), html.AY1("13.51"), html.AY2("17.49")),
		html.SvgLine(html.AX1("15.41"), html.AX2("8.59"), html.AY1("6.51"), html.AY2("10.49")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
