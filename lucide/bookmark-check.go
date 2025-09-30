package lucide

import (
	html "github.com/plainkit/html"
)

// BookmarkCheck creates a Bookmark Check Lucide icon.
func BookmarkCheck(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bookmark-check", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m19 21-7-4-7 4V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2Z")),
		html.SvgPath(html.AD("m9 10 2 2 4-4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
