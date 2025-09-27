package lucide

import (
	html "github.com/plainkit/html"
)

// MessageSquareText creates a Message Square Text Lucide icon.
func MessageSquareText(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-message-square-text", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M22 17a2 2 0 0 1-2 2H6.828a2 2 0 0 0-1.414.586l-2.202 2.202A.71.71 0 0 1 2 21.286V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2z"))),
		html.Child(html.SvgPath(html.AD("M7 11h10"))),
		html.Child(html.SvgPath(html.AD("M7 15h6"))),
		html.Child(html.SvgPath(html.AD("M7 7h8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
