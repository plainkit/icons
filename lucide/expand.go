package lucide

import (
	html "github.com/plainkit/html"
)

// Expand creates a Expand Lucide icon.
func Expand(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-expand", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m15 15 6 6")),
		html.SvgPath(html.AD("m15 9 6-6")),
		html.SvgPath(html.AD("M21 16v5h-5")),
		html.SvgPath(html.AD("M21 8V3h-5")),
		html.SvgPath(html.AD("M3 16v5h5")),
		html.SvgPath(html.AD("m3 21 6-6")),
		html.SvgPath(html.AD("M3 8V3h5")),
		html.SvgPath(html.AD("M9 9 3 3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
