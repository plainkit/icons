package lucide

import (
	html "github.com/plainkit/html"
)

// MessageCircle creates a Message Circle Lucide icon.
func MessageCircle(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-message-circle", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M2.992 16.342a2 2 0 0 1 .094 1.167l-1.065 3.29a1 1 0 0 0 1.236 1.168l3.413-.998a2 2 0 0 1 1.099.092 10 10 0 1 0-4.777-4.719"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
