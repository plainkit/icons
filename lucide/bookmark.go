package lucide

import (
	html "github.com/plainkit/html"
)

// Bookmark creates a Bookmark Lucide icon.
func Bookmark(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bookmark", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m19 21-7-4-7 4V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v16z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
