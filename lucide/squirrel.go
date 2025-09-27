package lucide

import (
	html "github.com/plainkit/html"
)

// Squirrel creates a Squirrel Lucide icon.
func Squirrel(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-squirrel", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M15.236 22a3 3 0 0 0-2.2-5"))),
		html.Child(html.SvgPath(html.AD("M16 20a3 3 0 0 1 3-3h1a2 2 0 0 0 2-2v-2a4 4 0 0 0-4-4V4"))),
		html.Child(html.SvgPath(html.AD("M18 13h.01"))),
		html.Child(html.SvgPath(html.AD("M18 6a4 4 0 0 0-4 4 7 7 0 0 0-7 7c0-5 4-5 4-10.5a4.5 4.5 0 1 0-9 0 2.5 2.5 0 0 0 5 0C7 10 3 11 3 17c0 2.8 2.2 5 5 5h10"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
