package lucide

import (
	html "github.com/plainkit/html"
)

// ScissorsLineDashed creates a Scissors Line Dashed Lucide icon.
func ScissorsLineDashed(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-scissors-line-dashed", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M5.42 9.42 8 12")),
		html.SvgCircle(html.ACx("4"), html.ACy("8"), html.AR("2")),
		html.SvgPath(html.AD("m14 6-8.58 8.58")),
		html.SvgCircle(html.ACx("4"), html.ACy("16"), html.AR("2")),
		html.SvgPath(html.AD("M10.8 14.8 14 18")),
		html.SvgPath(html.AD("M16 12h-2")),
		html.SvgPath(html.AD("M22 12h-2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
