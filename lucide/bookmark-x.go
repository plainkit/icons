package lucide

import (
	html "github.com/plainkit/html"
)

// BookmarkX creates a Bookmark X Lucide icon.
func BookmarkX(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bookmark-x", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m19 21-7-4-7 4V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2Z"))),
		html.Child(html.SvgPath(html.AD("m14.5 7.5-5 5"))),
		html.Child(html.SvgPath(html.AD("m9.5 7.5 5 5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
