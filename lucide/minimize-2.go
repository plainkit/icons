package lucide

import (
	html "github.com/plainkit/html"
)

// Minimize2 creates a Minimize 2 Lucide icon.
func Minimize2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-minimize-2", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m14 10 7-7")),
		html.SvgPath(html.AD("M20 10h-6V4")),
		html.SvgPath(html.AD("m3 21 7-7")),
		html.SvgPath(html.AD("M4 14h6v6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
