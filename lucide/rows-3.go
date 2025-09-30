package lucide

import (
	html "github.com/plainkit/html"
)

// Rows3 creates a Rows 3 Lucide icon.
func Rows3(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-rows-3", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M21 9H3")),
		html.SvgPath(html.AD("M21 15H3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
