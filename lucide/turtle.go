package lucide

import (
	html "github.com/plainkit/html"
)

// Turtle creates a Turtle Lucide icon.
func Turtle(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-turtle", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m12 10 2 4v3a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-3a8 8 0 1 0-16 0v3a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-3l2-4h4Z"))),
		html.Child(html.SvgPath(html.AD("M4.82 7.9 8 10"))),
		html.Child(html.SvgPath(html.AD("M15.18 7.9 12 10"))),
		html.Child(html.SvgPath(html.AD("M16.93 10H20a2 2 0 0 1 0 4H2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
