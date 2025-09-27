package lucide

import (
	html "github.com/plainkit/html"
)

// Handbag creates a Handbag Lucide icon.
func Handbag(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-handbag", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M2.048 18.566A2 2 0 0 0 4 21h16a2 2 0 0 0 1.952-2.434l-2-9A2 2 0 0 0 18 8H6a2 2 0 0 0-1.952 1.566z"))),
		html.Child(html.SvgPath(html.AD("M8 11V6a4 4 0 0 1 8 0v5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
