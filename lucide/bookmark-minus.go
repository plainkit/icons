package lucide

import (
	html "github.com/plainkit/html"
)

// BookmarkMinus creates a Bookmark Minus Lucide icon.
func BookmarkMinus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bookmark-minus", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m19 21-7-4-7 4V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v16z")),
		html.SvgLine(html.AX1("15"), html.AX2("9"), html.AY1("10"), html.AY2("10")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
