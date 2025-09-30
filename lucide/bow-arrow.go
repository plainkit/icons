package lucide

import (
	html "github.com/plainkit/html"
)

// BowArrow creates a Bow Arrow Lucide icon.
func BowArrow(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bow-arrow", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M17 3h4v4")),
		html.SvgPath(html.AD("M18.575 11.082a13 13 0 0 1 1.048 9.027 1.17 1.17 0 0 1-1.914.597L14 17")),
		html.SvgPath(html.AD("M7 10 3.29 6.29a1.17 1.17 0 0 1 .6-1.91 13 13 0 0 1 9.03 1.05")),
		html.SvgPath(html.AD("M7 14a1.7 1.7 0 0 0-1.207.5l-2.646 2.646A.5.5 0 0 0 3.5 18H5a1 1 0 0 1 1 1v1.5a.5.5 0 0 0 .854.354L9.5 18.207A1.7 1.7 0 0 0 10 17v-2a1 1 0 0 0-1-1z")),
		html.SvgPath(html.AD("M9.707 14.293 21 3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
