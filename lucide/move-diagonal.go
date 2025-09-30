package lucide

import (
	html "github.com/plainkit/html"
)

// MoveDiagonal creates a Move Diagonal Lucide icon.
func MoveDiagonal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-move-diagonal", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M11 19H5v-6")),
		html.SvgPath(html.AD("M13 5h6v6")),
		html.SvgPath(html.AD("M19 5 5 19")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
