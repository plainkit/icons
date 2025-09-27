package lucide

import (
	html "github.com/plainkit/html"
)

// ServerOff creates a Server Off Lucide icon.
func ServerOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-server-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M7 2h13a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-5"))),
		html.Child(html.SvgPath(html.AD("M10 10 2.5 2.5C2 2 2 2.5 2 5v3a2 2 0 0 0 2 2h6z"))),
		html.Child(html.SvgPath(html.AD("M22 17v-1a2 2 0 0 0-2-2h-1"))),
		html.Child(html.SvgPath(html.AD("M4 14a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h16.5l1-.5.5.5-8-8H4z"))),
		html.Child(html.SvgPath(html.AD("M6 18h.01"))),
		html.Child(html.SvgPath(html.AD("m2 2 20 20"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
