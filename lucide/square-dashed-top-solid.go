package lucide

import (
	html "github.com/plainkit/html"
)

// SquareDashedTopSolid creates a Square Dashed Top Solid Lucide icon.
func SquareDashedTopSolid(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-dashed-top-solid", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M14 21h1")),
		html.SvgPath(html.AD("M21 14v1")),
		html.SvgPath(html.AD("M21 19a2 2 0 0 1-2 2")),
		html.SvgPath(html.AD("M21 9v1")),
		html.SvgPath(html.AD("M3 14v1")),
		html.SvgPath(html.AD("M3 5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2")),
		html.SvgPath(html.AD("M3 9v1")),
		html.SvgPath(html.AD("M5 21a2 2 0 0 1-2-2")),
		html.SvgPath(html.AD("M9 21h1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
