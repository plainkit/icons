package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowDownLeft creates a Arrow Down Left Lucide icon.
func ArrowDownLeft(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-down-left", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M17 7 7 17"))),
		html.Child(html.SvgPath(html.AD("M17 17H7V7"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
