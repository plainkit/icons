package lucide

import (
	html "github.com/plainkit/html"
)

// RotateCcwSquare creates a Rotate Ccw Square Lucide icon.
func RotateCcwSquare(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-rotate-ccw-square", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M20 9V7a2 2 0 0 0-2-2h-6")),
		html.SvgPath(html.AD("m15 2-3 3 3 3")),
		html.SvgPath(html.AD("M20 13v5a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
