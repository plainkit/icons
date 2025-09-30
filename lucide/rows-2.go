package lucide

import (
	html "github.com/plainkit/html"
)

// Rows2 creates a Rows 2 Lucide icon.
func Rows2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-rows-2", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M3 12h18")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
