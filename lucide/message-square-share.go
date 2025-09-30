package lucide

import (
	html "github.com/plainkit/html"
)

// MessageSquareShare creates a Message Square Share Lucide icon.
func MessageSquareShare(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-message-square-share", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 3H4a2 2 0 0 0-2 2v16.286a.71.71 0 0 0 1.212.502l2.202-2.202A2 2 0 0 1 6.828 19H20a2 2 0 0 0 2-2v-4")),
		html.SvgPath(html.AD("M16 3h6v6")),
		html.SvgPath(html.AD("m16 9 6-6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
