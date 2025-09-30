package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowUpToLine creates a Arrow Up To Line Lucide icon.
func ArrowUpToLine(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-up-to-line", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M5 3h14")),
		html.SvgPath(html.AD("m18 13-6-6-6 6")),
		html.SvgPath(html.AD("M12 7v14")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
