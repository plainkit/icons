package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowDownToLine creates a Arrow Down To Line Lucide icon.
func ArrowDownToLine(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-down-to-line", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 17V3"))),
		html.Child(html.SvgPath(html.AD("m6 11 6 6 6-6"))),
		html.Child(html.SvgPath(html.AD("M19 21H5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
