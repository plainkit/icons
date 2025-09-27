package lucide

import (
	html "github.com/plainkit/html"
)

// Snowflake creates a Snowflake Lucide icon.
func Snowflake(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-snowflake", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m10 20-1.25-2.5L6 18"))),
		html.Child(html.SvgPath(html.AD("M10 4 8.75 6.5 6 6"))),
		html.Child(html.SvgPath(html.AD("m14 20 1.25-2.5L18 18"))),
		html.Child(html.SvgPath(html.AD("m14 4 1.25 2.5L18 6"))),
		html.Child(html.SvgPath(html.AD("m17 21-3-6h-4"))),
		html.Child(html.SvgPath(html.AD("m17 3-3 6 1.5 3"))),
		html.Child(html.SvgPath(html.AD("M2 12h6.5L10 9"))),
		html.Child(html.SvgPath(html.AD("m20 10-1.5 2 1.5 2"))),
		html.Child(html.SvgPath(html.AD("M22 12h-6.5L14 15"))),
		html.Child(html.SvgPath(html.AD("m4 10 1.5 2L4 14"))),
		html.Child(html.SvgPath(html.AD("m7 21 3-6-1.5-3"))),
		html.Child(html.SvgPath(html.AD("m7 3 3 6h4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
