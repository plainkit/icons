package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowDownNarrowWide creates a Arrow Down Narrow Wide Lucide icon.
func ArrowDownNarrowWide(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-down-narrow-wide", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m3 16 4 4 4-4")),
		html.SvgPath(html.AD("M7 20V4")),
		html.SvgPath(html.AD("M11 4h4")),
		html.SvgPath(html.AD("M11 8h7")),
		html.SvgPath(html.AD("M11 12h10")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
