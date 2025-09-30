package lucide

import (
	html "github.com/plainkit/html"
)

// Maximize creates a Maximize Lucide icon.
func Maximize(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-maximize", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M8 3H5a2 2 0 0 0-2 2v3")),
		html.SvgPath(html.AD("M21 8V5a2 2 0 0 0-2-2h-3")),
		html.SvgPath(html.AD("M3 16v3a2 2 0 0 0 2 2h3")),
		html.SvgPath(html.AD("M16 21h3a2 2 0 0 0 2-2v-3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
