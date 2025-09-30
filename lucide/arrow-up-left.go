package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowUpLeft creates a Arrow Up Left Lucide icon.
func ArrowUpLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-up-left", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M7 17V7h10")),
		html.SvgPath(html.AD("M17 17 7 7")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
