package lucide

import (
	html "github.com/plainkit/html"
)

// BellOff creates a Bell Off Lucide icon.
func BellOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bell-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10.268 21a2 2 0 0 0 3.464 0"))),
		html.Child(html.SvgPath(html.AD("M17 17H4a1 1 0 0 1-.74-1.673C4.59 13.956 6 12.499 6 8a6 6 0 0 1 .258-1.742"))),
		html.Child(html.SvgPath(html.AD("m2 2 20 20"))),
		html.Child(html.SvgPath(html.AD("M8.668 3.01A6 6 0 0 1 18 8c0 2.687.77 4.653 1.707 6.05"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
