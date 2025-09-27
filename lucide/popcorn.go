package lucide

import (
	html "github.com/plainkit/html"
)

// Popcorn creates a Popcorn Lucide icon.
func Popcorn(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-popcorn", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M18 8a2 2 0 0 0 0-4 2 2 0 0 0-4 0 2 2 0 0 0-4 0 2 2 0 0 0-4 0 2 2 0 0 0 0 4"))),
		html.Child(html.SvgPath(html.AD("M10 22 9 8"))),
		html.Child(html.SvgPath(html.AD("m14 22 1-14"))),
		html.Child(html.SvgPath(html.AD("M20 8c.5 0 .9.4.8 1l-2.6 12c-.1.5-.7 1-1.2 1H7c-.6 0-1.1-.4-1.2-1L3.2 9c-.1-.6.3-1 .8-1Z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
