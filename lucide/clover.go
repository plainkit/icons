package lucide

import (
	html "github.com/plainkit/html"
)

// Clover creates a Clover Lucide icon.
func Clover(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-clover", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M16.17 7.83 2 22"))),
		html.Child(html.SvgPath(html.AD("M4.02 12a2.827 2.827 0 1 1 3.81-4.17A2.827 2.827 0 1 1 12 4.02a2.827 2.827 0 1 1 4.17 3.81A2.827 2.827 0 1 1 19.98 12a2.827 2.827 0 1 1-3.81 4.17A2.827 2.827 0 1 1 12 19.98a2.827 2.827 0 1 1-4.17-3.81A1 1 0 1 1 4 12"))),
		html.Child(html.SvgPath(html.AD("m7.83 7.83 8.34 8.34"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
