package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowLeftFromLine creates a Arrow Left From Line Lucide icon.
func ArrowLeftFromLine(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-left-from-line", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m9 6-6 6 6 6")),
		html.SvgPath(html.AD("M3 12h14")),
		html.SvgPath(html.AD("M21 19V5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
