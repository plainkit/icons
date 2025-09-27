package lucide

import (
	html "github.com/plainkit/html"
)

// Tent creates a Tent Lucide icon.
func Tent(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-tent", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3.5 21 14 3"))),
		html.Child(html.SvgPath(html.AD("M20.5 21 10 3"))),
		html.Child(html.SvgPath(html.AD("M15.5 21 12 15l-3.5 6"))),
		html.Child(html.SvgPath(html.AD("M2 21h20"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
