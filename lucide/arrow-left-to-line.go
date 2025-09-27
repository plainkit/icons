package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowLeftToLine creates a Arrow Left To Line Lucide icon.
func ArrowLeftToLine(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-left-to-line", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3 19V5"))),
		html.Child(html.SvgPath(html.AD("m13 6-6 6 6 6"))),
		html.Child(html.SvgPath(html.AD("M7 12h14"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
