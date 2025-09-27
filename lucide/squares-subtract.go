package lucide

import (
	html "github.com/plainkit/html"
)

// SquaresSubtract creates a Squares Subtract Lucide icon.
func SquaresSubtract(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-squares-subtract", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 22a2 2 0 0 1-2-2"))),
		html.Child(html.SvgPath(html.AD("M16 22h-2"))),
		html.Child(html.SvgPath(html.AD("M16 4a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h3a1 1 0 0 0 1-1v-5a2 2 0 0 1 2-2h5a1 1 0 0 0 1-1z"))),
		html.Child(html.SvgPath(html.AD("M20 8a2 2 0 0 1 2 2"))),
		html.Child(html.SvgPath(html.AD("M22 14v2"))),
		html.Child(html.SvgPath(html.AD("M22 20a2 2 0 0 1-2 2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
