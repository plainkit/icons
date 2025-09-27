package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowDownUp creates a Arrow Down Up Lucide icon.
func ArrowDownUp(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-down-up", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m3 16 4 4 4-4"))),
		html.Child(html.SvgPath(html.AD("M7 20V4"))),
		html.Child(html.SvgPath(html.AD("m21 8-4-4-4 4"))),
		html.Child(html.SvgPath(html.AD("M17 4v16"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
