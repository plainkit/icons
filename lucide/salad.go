package lucide

import (
	html "github.com/plainkit/html"
)

// Salad creates a Salad Lucide icon.
func Salad(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-salad", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M7 21h10"))),
		html.Child(html.SvgPath(html.AD("M12 21a9 9 0 0 0 9-9H3a9 9 0 0 0 9 9Z"))),
		html.Child(html.SvgPath(html.AD("M11.38 12a2.4 2.4 0 0 1-.4-4.77 2.4 2.4 0 0 1 3.2-2.77 2.4 2.4 0 0 1 3.47-.63 2.4 2.4 0 0 1 3.37 3.37 2.4 2.4 0 0 1-1.1 3.7 2.51 2.51 0 0 1 .03 1.1"))),
		html.Child(html.SvgPath(html.AD("m13 12 4-4"))),
		html.Child(html.SvgPath(html.AD("M10.9 7.25A3.99 3.99 0 0 0 4 10c0 .73.2 1.41.54 2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
