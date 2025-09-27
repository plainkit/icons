package lucide

import (
	html "github.com/plainkit/html"
)

// BookText creates a Book Text Lucide icon.
func BookText(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-book-text", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"))),
		html.Child(html.SvgPath(html.AD("M8 11h8"))),
		html.Child(html.SvgPath(html.AD("M8 7h6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
