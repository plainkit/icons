package lucide

import (
	html "github.com/plainkit/html"
)

// BadgeIndianRupee creates a Badge Indian Rupee Lucide icon.
func BadgeIndianRupee(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-badge-indian-rupee", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z")),
		html.SvgPath(html.AD("M8 8h8")),
		html.SvgPath(html.AD("M8 12h8")),
		html.SvgPath(html.AD("m13 17-5-1h1a4 4 0 0 0 0-8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
