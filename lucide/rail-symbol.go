package lucide

import (
	html "github.com/plainkit/html"
)

// RailSymbol creates a Rail Symbol Lucide icon.
func RailSymbol(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-rail-symbol", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M5 15h14")),
		html.SvgPath(html.AD("M5 9h14")),
		html.SvgPath(html.AD("m14 20-5-5 6-6-5-5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
