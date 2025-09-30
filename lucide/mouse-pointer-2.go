package lucide

import (
	html "github.com/plainkit/html"
)

// MousePointer2 creates a Mouse Pointer 2 Lucide icon.
func MousePointer2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-mouse-pointer-2", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4.037 4.688a.495.495 0 0 1 .651-.651l16 6.5a.5.5 0 0 1-.063.947l-6.124 1.58a2 2 0 0 0-1.438 1.435l-1.579 6.126a.5.5 0 0 1-.947.063z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
