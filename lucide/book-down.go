package lucide

import (
	html "github.com/plainkit/html"
)

// BookDown creates a Book Down Lucide icon.
func BookDown(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-book-down", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 13V7"))),
		html.Child(html.SvgPath(html.AD("M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"))),
		html.Child(html.SvgPath(html.AD("m9 10 3 3 3-3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
