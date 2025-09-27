package lucide

import (
	html "github.com/plainkit/html"
)

// Rows4 creates a Rows 4 Lucide icon.
func Rows4(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-rows-4", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M21 7.5H3"))),
		html.Child(html.SvgPath(html.AD("M21 12H3"))),
		html.Child(html.SvgPath(html.AD("M21 16.5H3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
