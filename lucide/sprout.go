package lucide

import (
	html "github.com/plainkit/html"
)

// Sprout creates a Sprout Lucide icon.
func Sprout(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-sprout", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M14 9.536V7a4 4 0 0 1 4-4h1.5a.5.5 0 0 1 .5.5V5a4 4 0 0 1-4 4 4 4 0 0 0-4 4c0 2 1 3 1 5a5 5 0 0 1-1 3")),
		html.SvgPath(html.AD("M4 9a5 5 0 0 1 8 4 5 5 0 0 1-8-4")),
		html.SvgPath(html.AD("M5 21h14")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
