package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowDownFromLine creates a Arrow Down From Line Lucide icon.
func ArrowDownFromLine(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-down-from-line", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M19 3H5")),
		html.SvgPath(html.AD("M12 21V7")),
		html.SvgPath(html.AD("m6 15 6 6 6-6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
