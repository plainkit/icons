package lucide

import (
	html "github.com/plainkit/html"
)

// SeparatorHorizontal creates a Separator Horizontal Lucide icon.
func SeparatorHorizontal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-separator-horizontal", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m16 16-4 4-4-4")),
		html.SvgPath(html.AD("M3 12h18")),
		html.SvgPath(html.AD("m8 8 4-4 4 4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
