package lucide

import (
	html "github.com/plainkit/html"
)

// RedoDot creates a Redo Dot Lucide icon.
func RedoDot(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-redo-dot", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("17"), html.AR("1")),
		html.SvgPath(html.AD("M21 7v6h-6")),
		html.SvgPath(html.AD("M3 17a9 9 0 0 1 9-9 9 9 0 0 1 6 2.3l3 2.7")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
