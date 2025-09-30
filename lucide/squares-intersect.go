package lucide

import (
	html "github.com/plainkit/html"
)

// SquaresIntersect creates a Squares Intersect Lucide icon.
func SquaresIntersect(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-squares-intersect", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 22a2 2 0 0 1-2-2")),
		html.SvgPath(html.AD("M14 2a2 2 0 0 1 2 2")),
		html.SvgPath(html.AD("M16 22h-2")),
		html.SvgPath(html.AD("M2 10V8")),
		html.SvgPath(html.AD("M2 4a2 2 0 0 1 2-2")),
		html.SvgPath(html.AD("M20 8a2 2 0 0 1 2 2")),
		html.SvgPath(html.AD("M22 14v2")),
		html.SvgPath(html.AD("M22 20a2 2 0 0 1-2 2")),
		html.SvgPath(html.AD("M4 16a2 2 0 0 1-2-2")),
		html.SvgPath(html.AD("M8 10a2 2 0 0 1 2-2h5a1 1 0 0 1 1 1v5a2 2 0 0 1-2 2H9a1 1 0 0 1-1-1z")),
		html.SvgPath(html.AD("M8 2h2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
