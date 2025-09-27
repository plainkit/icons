package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowUpNarrowWide creates a Arrow Up Narrow Wide Lucide icon.
func ArrowUpNarrowWide(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-up-narrow-wide", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m3 8 4-4 4 4"))),
		html.Child(html.SvgPath(html.AD("M7 4v16"))),
		html.Child(html.SvgPath(html.AD("M11 12h4"))),
		html.Child(html.SvgPath(html.AD("M11 16h7"))),
		html.Child(html.SvgPath(html.AD("M11 20h10"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
