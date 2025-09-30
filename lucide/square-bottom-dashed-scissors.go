package lucide

import (
	html "github.com/plainkit/html"
)

// SquareBottomDashedScissors creates a Square Bottom Dashed Scissors Lucide icon.
func SquareBottomDashedScissors(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-bottom-dashed-scissors", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2")),
		html.SvgPath(html.AD("M10 22H8")),
		html.SvgPath(html.AD("M16 22h-2")),
		html.SvgCircle(html.ACx("8"), html.ACy("8"), html.AR("2")),
		html.SvgPath(html.AD("M9.414 9.414 12 12")),
		html.SvgPath(html.AD("M14.8 14.8 18 18")),
		html.SvgCircle(html.ACx("8"), html.ACy("16"), html.AR("2")),
		html.SvgPath(html.AD("m18 6-8.586 8.586")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
