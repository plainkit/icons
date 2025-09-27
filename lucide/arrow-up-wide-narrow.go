package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowUpWideNarrow creates a Arrow Up Wide Narrow Lucide icon.
func ArrowUpWideNarrow(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-up-wide-narrow", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m3 8 4-4 4 4"))),
		html.Child(html.SvgPath(html.AD("M7 4v16"))),
		html.Child(html.SvgPath(html.AD("M11 12h10"))),
		html.Child(html.SvgPath(html.AD("M11 16h7"))),
		html.Child(html.SvgPath(html.AD("M11 20h4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
