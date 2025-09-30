package lucide

import (
	html "github.com/plainkit/html"
)

// Cigarette creates a Cigarette Lucide icon.
func Cigarette(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cigarette", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M17 12H3a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h14")),
		html.SvgPath(html.AD("M18 8c0-2.5-2-2.5-2-5")),
		html.SvgPath(html.AD("M21 16a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1")),
		html.SvgPath(html.AD("M22 8c0-2.5-2-2.5-2-5")),
		html.SvgPath(html.AD("M7 12v4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
