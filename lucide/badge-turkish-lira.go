package lucide

import (
	html "github.com/plainkit/html"
)

// BadgeTurkishLira creates a Badge Turkish Lira Lucide icon.
func BadgeTurkishLira(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-badge-turkish-lira", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M11 7v10a5 5 0 0 0 5-5"))),
		html.Child(html.SvgPath(html.AD("m15 8-6 3"))),
		html.Child(html.SvgPath(html.AD("M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
