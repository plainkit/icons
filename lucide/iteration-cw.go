package lucide

import (
	html "github.com/plainkit/html"
)

// IterationCw creates a Iteration Cw Lucide icon.
func IterationCw(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-iteration-cw", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M4 10a8 8 0 1 1 8 8H4"))),
		html.Child(html.SvgPath(html.AD("m8 22-4-4 4-4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
