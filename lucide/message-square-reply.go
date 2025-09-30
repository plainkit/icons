package lucide

import (
	html "github.com/plainkit/html"
)

// MessageSquareReply creates a Message Square Reply Lucide icon.
func MessageSquareReply(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-message-square-reply", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M22 17a2 2 0 0 1-2 2H6.828a2 2 0 0 0-1.414.586l-2.202 2.202A.71.71 0 0 1 2 21.286V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2z")),
		html.SvgPath(html.AD("m10 8-3 3 3 3")),
		html.SvgPath(html.AD("M17 14v-1a2 2 0 0 0-2-2H7")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
