package lucide

import (
	html "github.com/plainkit/html"
)

// RotateCwSquare creates a Rotate Cw Square Lucide icon.
func RotateCwSquare(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-rotate-cw-square", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 5H6a2 2 0 0 0-2 2v3")),
		html.SvgPath(html.AD("m9 8 3-3-3-3")),
		html.SvgPath(html.AD("M4 14v4a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2h-2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
