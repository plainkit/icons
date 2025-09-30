package lucide

import (
	html "github.com/plainkit/html"
)

// BookCheck creates a Book Check Lucide icon.
func BookCheck(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-book-check", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20")),
		html.SvgPath(html.AD("m9 9.5 2 2 4-4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
