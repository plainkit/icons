package lucide

import (
	html "github.com/plainkit/html"
)

// BookMinus creates a Book Minus Lucide icon.
func BookMinus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-book-minus", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20")),
		html.SvgPath(html.AD("M9 10h6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
