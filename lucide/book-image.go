package lucide

import (
	html "github.com/plainkit/html"
)

// BookImage creates a Book Image Lucide icon.
func BookImage(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-book-image", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m20 13.7-2.1-2.1a2 2 0 0 0-2.8 0L9.7 17")),
		html.SvgPath(html.AD("M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H19a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6.5a1 1 0 0 1 0-5H20")),
		html.SvgCircle(html.ACx("10"), html.ACy("8"), html.AR("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
