package lucide

import (
	html "github.com/plainkit/html"
)

// Combine creates a Combine Lucide icon.
func Combine(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-combine", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 18H5a3 3 0 0 1-3-3v-1")),
		html.SvgPath(html.AD("M14 2a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2")),
		html.SvgPath(html.AD("M20 2a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2")),
		html.SvgPath(html.AD("m7 21 3-3-3-3")),
		html.SvgRect(html.AWidth("8"), html.AHeight("8"), html.AX("14"), html.AY("14"), html.ARx("2")),
		html.SvgRect(html.AWidth("8"), html.AHeight("8"), html.AX("2"), html.AY("2"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
