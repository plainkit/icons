package lucide

import (
	html "github.com/plainkit/html"
)

// Minimize creates a Minimize Lucide icon.
func Minimize(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-minimize", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M8 3v3a2 2 0 0 1-2 2H3")),
		html.SvgPath(html.AD("M21 8h-3a2 2 0 0 1-2-2V3")),
		html.SvgPath(html.AD("M3 16h3a2 2 0 0 1 2 2v3")),
		html.SvgPath(html.AD("M16 21v-3a2 2 0 0 1 2-2h3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
