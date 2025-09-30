package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowRightToLine creates a Arrow Right To Line Lucide icon.
func ArrowRightToLine(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-right-to-line", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M17 12H3")),
		html.SvgPath(html.AD("m11 18 6-6-6-6")),
		html.SvgPath(html.AD("M21 5v14")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
