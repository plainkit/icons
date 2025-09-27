package lucide

import (
	html "github.com/plainkit/html"
)

// Refrigerator creates a Refrigerator Lucide icon.
func Refrigerator(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-refrigerator", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M5 6a4 4 0 0 1 4-4h6a4 4 0 0 1 4 4v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6Z"))),
		html.Child(html.SvgPath(html.AD("M5 10h14"))),
		html.Child(html.SvgPath(html.AD("M15 7v6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
