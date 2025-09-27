package lucide

import (
	html "github.com/plainkit/html"
)

// MessageSquareDiff creates a Message Square Diff Lucide icon.
func MessageSquareDiff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-message-square-diff", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M22 17a2 2 0 0 1-2 2H6.828a2 2 0 0 0-1.414.586l-2.202 2.202A.71.71 0 0 1 2 21.286V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2z"))),
		html.Child(html.SvgPath(html.AD("M10 15h4"))),
		html.Child(html.SvgPath(html.AD("M10 9h4"))),
		html.Child(html.SvgPath(html.AD("M12 7v4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
