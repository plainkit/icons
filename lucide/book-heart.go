package lucide

import (
	html "github.com/plainkit/html"
)

// BookHeart creates a Book Heart Lucide icon.
func BookHeart(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-book-heart", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20"))),
		html.Child(html.SvgPath(html.AD("M8.62 9.8A2.25 2.25 0 1 1 12 6.836a2.25 2.25 0 1 1 3.38 2.966l-2.626 2.856a.998.998 0 0 1-1.507 0z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
