package lucide

import (
	html "github.com/plainkit/html"
)

// Lightbulb creates a Lightbulb Lucide icon.
func Lightbulb(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-lightbulb", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M15 14c.2-1 .7-1.7 1.5-2.5 1-.9 1.5-2.2 1.5-3.5A6 6 0 0 0 6 8c0 1 .2 2.2 1.5 3.5.7.7 1.3 1.5 1.5 2.5"))),
		html.Child(html.SvgPath(html.AD("M9 18h6"))),
		html.Child(html.SvgPath(html.AD("M10 22h4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
