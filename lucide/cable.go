package lucide

import (
	html "github.com/plainkit/html"
)

// Cable creates a Cable Lucide icon.
func Cable(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cable", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M17 19a1 1 0 0 1-1-1v-2a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2a1 1 0 0 1-1 1z"))),
		html.Child(html.SvgPath(html.AD("M17 21v-2"))),
		html.Child(html.SvgPath(html.AD("M19 14V6.5a1 1 0 0 0-7 0v11a1 1 0 0 1-7 0V10"))),
		html.Child(html.SvgPath(html.AD("M21 21v-2"))),
		html.Child(html.SvgPath(html.AD("M3 5V3"))),
		html.Child(html.SvgPath(html.AD("M4 10a2 2 0 0 1-2-2V6a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v2a2 2 0 0 1-2 2z"))),
		html.Child(html.SvgPath(html.AD("M7 5V3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
