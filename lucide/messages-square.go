package lucide

import (
	html "github.com/plainkit/html"
)

// MessagesSquare creates a Messages Square Lucide icon.
func MessagesSquare(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-messages-square", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M16 10a2 2 0 0 1-2 2H6.828a2 2 0 0 0-1.414.586l-2.202 2.202A.71.71 0 0 1 2 14.286V4a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z")),
		html.SvgPath(html.AD("M20 9a2 2 0 0 1 2 2v10.286a.71.71 0 0 1-1.212.502l-2.202-2.202A2 2 0 0 0 17.172 19H10a2 2 0 0 1-2-2v-1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
