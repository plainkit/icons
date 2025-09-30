package lucide

import (
	html "github.com/plainkit/html"
)

// SquareDashedBottomCode creates a Square Dashed Bottom Code Lucide icon.
func SquareDashedBottomCode(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-dashed-bottom-code", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 9.5 8 12l2 2.5")),
		html.SvgPath(html.AD("M14 21h1")),
		html.SvgPath(html.AD("m14 9.5 2 2.5-2 2.5")),
		html.SvgPath(html.AD("M5 21a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2")),
		html.SvgPath(html.AD("M9 21h1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
