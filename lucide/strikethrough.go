package lucide

import (
	html "github.com/plainkit/html"
)

// Strikethrough creates a Strikethrough Lucide icon.
func Strikethrough(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-strikethrough", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M16 4H9a3 3 0 0 0-2.83 4")),
		html.SvgPath(html.AD("M14 12a4 4 0 0 1 0 8H6")),
		html.SvgLine(html.AX1("4"), html.AX2("20"), html.AY1("12"), html.AY2("12")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
