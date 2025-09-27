package lucide

import (
	html "github.com/plainkit/html"
)

// Handshake creates a Handshake Lucide icon.
func Handshake(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-handshake", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m11 17 2 2a1 1 0 1 0 3-3"))),
		html.Child(html.SvgPath(html.AD("m14 14 2.5 2.5a1 1 0 1 0 3-3l-3.88-3.88a3 3 0 0 0-4.24 0l-.88.88a1 1 0 1 1-3-3l2.81-2.81a5.79 5.79 0 0 1 7.06-.87l.47.28a2 2 0 0 0 1.42.25L21 4"))),
		html.Child(html.SvgPath(html.AD("m21 3 1 11h-2"))),
		html.Child(html.SvgPath(html.AD("M3 3 2 14l6.5 6.5a1 1 0 1 0 3-3"))),
		html.Child(html.SvgPath(html.AD("M3 4h8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
