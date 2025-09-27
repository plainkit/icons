package lucide

import (
	html "github.com/plainkit/html"
)

// ChevronsUpDown creates a Chevrons Up Down Lucide icon.
func ChevronsUpDown(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chevrons-up-down", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m7 15 5 5 5-5"))),
		html.Child(html.SvgPath(html.AD("m7 9 5-5 5 5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
