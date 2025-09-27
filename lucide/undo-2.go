package lucide

import (
	html "github.com/plainkit/html"
)

// Undo2 creates a Undo 2 Lucide icon.
func Undo2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-undo-2", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M9 14 4 9l5-5"))),
		html.Child(html.SvgPath(html.AD("M4 9h10.5a5.5 5.5 0 0 1 5.5 5.5a5.5 5.5 0 0 1-5.5 5.5H11"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
