package lucide

import (
	html "github.com/plainkit/html"
)

// Theater creates a Theater Lucide icon.
func Theater(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-theater", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M2 10s3-3 3-8"))),
		html.Child(html.SvgPath(html.AD("M22 10s-3-3-3-8"))),
		html.Child(html.SvgPath(html.AD("M10 2c0 4.4-3.6 8-8 8"))),
		html.Child(html.SvgPath(html.AD("M14 2c0 4.4 3.6 8 8 8"))),
		html.Child(html.SvgPath(html.AD("M2 10s2 2 2 5"))),
		html.Child(html.SvgPath(html.AD("M22 10s-2 2-2 5"))),
		html.Child(html.SvgPath(html.AD("M8 15h8"))),
		html.Child(html.SvgPath(html.AD("M2 22v-1a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v1"))),
		html.Child(html.SvgPath(html.AD("M14 22v-1a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
