package lucide

import (
	html "github.com/plainkit/html"
)

// Binary creates a Binary Lucide icon.
func Binary(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-binary", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("4"), html.AHeight("6"), html.AX("14"), html.AY("14"), html.ARx("2")),
		html.SvgRect(html.AWidth("4"), html.AHeight("6"), html.AX("6"), html.AY("4"), html.ARx("2")),
		html.SvgPath(html.AD("M6 20h4")),
		html.SvgPath(html.AD("M14 10h4")),
		html.SvgPath(html.AD("M6 14h2v6")),
		html.SvgPath(html.AD("M14 4h2v6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
