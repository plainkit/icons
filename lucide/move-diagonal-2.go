package lucide

import (
	html "github.com/plainkit/html"
)

// MoveDiagonal2 creates a Move Diagonal 2 Lucide icon.
func MoveDiagonal2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-move-diagonal-2", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M19 13v6h-6")),
		html.SvgPath(html.AD("M5 11V5h6")),
		html.SvgPath(html.AD("m5 5 14 14")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
