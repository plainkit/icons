package lucide

import (
	html "github.com/plainkit/html"
)

// MessageSquareCode creates a Message Square Code Lucide icon.
func MessageSquareCode(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-message-square-code", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M22 17a2 2 0 0 1-2 2H6.828a2 2 0 0 0-1.414.586l-2.202 2.202A.71.71 0 0 1 2 21.286V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2z"))),
		html.Child(html.SvgPath(html.AD("m10 8-3 3 3 3"))),
		html.Child(html.SvgPath(html.AD("m14 14 3-3-3-3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
