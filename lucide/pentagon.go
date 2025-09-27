package lucide

import (
	html "github.com/plainkit/html"
)

// Pentagon creates a Pentagon Lucide icon.
func Pentagon(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-pentagon", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10.83 2.38a2 2 0 0 1 2.34 0l8 5.74a2 2 0 0 1 .73 2.25l-3.04 9.26a2 2 0 0 1-1.9 1.37H7.04a2 2 0 0 1-1.9-1.37L2.1 10.37a2 2 0 0 1 .73-2.25z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
