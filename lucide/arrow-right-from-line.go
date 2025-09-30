package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowRightFromLine creates a Arrow Right From Line Lucide icon.
func ArrowRightFromLine(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-right-from-line", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M3 5v14")),
		html.SvgPath(html.AD("M21 12H7")),
		html.SvgPath(html.AD("m15 18 6-6-6-6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
