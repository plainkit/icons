package lucide

import (
	html "github.com/plainkit/html"
)

// X creates a X Lucide icon.
func X(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-x", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M18 6 6 18"))),
		html.Child(html.SvgPath(html.AD("m6 6 12 12"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
