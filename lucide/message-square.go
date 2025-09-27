package lucide

import (
	html "github.com/plainkit/html"
)

// MessageSquare creates a Message Square Lucide icon.
func MessageSquare(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-message-square", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M22 17a2 2 0 0 1-2 2H6.828a2 2 0 0 0-1.414.586l-2.202 2.202A.71.71 0 0 1 2 21.286V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
