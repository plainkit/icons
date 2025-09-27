package lucide

import (
	html "github.com/plainkit/html"
)

// ServerCrash creates a Server Crash Lucide icon.
func ServerCrash(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-server-crash", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M6 10H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-2"))),
		html.Child(html.SvgPath(html.AD("M6 14H4a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2h-2"))),
		html.Child(html.SvgPath(html.AD("M6 6h.01"))),
		html.Child(html.SvgPath(html.AD("M6 18h.01"))),
		html.Child(html.SvgPath(html.AD("m13 6-4 6h6l-4 6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
