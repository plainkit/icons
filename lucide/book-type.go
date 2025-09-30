package lucide

import (
	html "github.com/plainkit/html"
)

// BookType creates a Book Type Lucide icon.
func BookType(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-book-type", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 13h4")),
		html.SvgPath(html.AD("M12 6v7")),
		html.SvgPath(html.AD("M16 8V6H8v2")),
		html.SvgPath(html.AD("M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
