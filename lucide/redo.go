package lucide

import (
	html "github.com/plainkit/html"
)

// Redo creates a Redo Lucide icon.
func Redo(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-redo", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M21 7v6h-6"))),
		html.Child(html.SvgPath(html.AD("M3 17a9 9 0 0 1 9-9 9 9 0 0 1 6 2.3l3 2.7"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
