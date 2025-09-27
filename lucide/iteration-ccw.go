package lucide

import (
	html "github.com/plainkit/html"
)

// IterationCcw creates a Iteration Ccw Lucide icon.
func IterationCcw(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-iteration-ccw", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m16 14 4 4-4 4"))),
		html.Child(html.SvgPath(html.AD("M20 10a8 8 0 1 0-8 8h8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
