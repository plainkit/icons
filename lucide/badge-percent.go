package lucide

import (
	html "github.com/plainkit/html"
)

// BadgePercent creates a Badge Percent Lucide icon.
func BadgePercent(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-badge-percent", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"))),
		html.Child(html.SvgPath(html.AD("m15 9-6 6"))),
		html.Child(html.SvgPath(html.AD("M9 9h.01"))),
		html.Child(html.SvgPath(html.AD("M15 15h.01"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
