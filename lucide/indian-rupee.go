package lucide

import (
	html "github.com/plainkit/html"
)

// IndianRupee creates a Indian Rupee Lucide icon.
func IndianRupee(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-indian-rupee", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M6 3h12")),
		html.SvgPath(html.AD("M6 8h12")),
		html.SvgPath(html.AD("m6 13 8.5 8")),
		html.SvgPath(html.AD("M6 13h3")),
		html.SvgPath(html.AD("M9 13c6.667 0 6.667-10 0-10")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
