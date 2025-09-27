package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowDownWideNarrow creates a Arrow Down Wide Narrow Lucide icon.
func ArrowDownWideNarrow(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-down-wide-narrow", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m3 16 4 4 4-4"))),
		html.Child(html.SvgPath(html.AD("M7 20V4"))),
		html.Child(html.SvgPath(html.AD("M11 4h10"))),
		html.Child(html.SvgPath(html.AD("M11 8h7"))),
		html.Child(html.SvgPath(html.AD("M11 12h4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
