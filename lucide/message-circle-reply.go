package lucide

import (
	html "github.com/plainkit/html"
)

// MessageCircleReply creates a Message Circle Reply Lucide icon.
func MessageCircleReply(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-message-circle-reply", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M2.992 16.342a2 2 0 0 1 .094 1.167l-1.065 3.29a1 1 0 0 0 1.236 1.168l3.413-.998a2 2 0 0 1 1.099.092 10 10 0 1 0-4.777-4.719"))),
		html.Child(html.SvgPath(html.AD("m10 15-3-3 3-3"))),
		html.Child(html.SvgPath(html.AD("M7 12h8a2 2 0 0 1 2 2v1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
