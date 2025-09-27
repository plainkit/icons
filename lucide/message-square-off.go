package lucide

import (
	html "github.com/plainkit/html"
)

// MessageSquareOff creates a Message Square Off Lucide icon.
func MessageSquareOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-message-square-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M19 19H6.828a2 2 0 0 0-1.414.586l-2.202 2.202A.7.7 0 0 1 2 21.286V5a2 2 0 0 1 1.184-1.826"))),
		html.Child(html.SvgPath(html.AD("m2 2 20 20"))),
		html.Child(html.SvgPath(html.AD("M8.656 3H20a2 2 0 0 1 2 2v11.344"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
