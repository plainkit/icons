package lucide

import (
	html "github.com/plainkit/html"
)

// ChevronsLeftRightEllipsis creates a Chevrons Left Right Ellipsis Lucide icon.
func ChevronsLeftRightEllipsis(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chevrons-left-right-ellipsis", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 12h.01"))),
		html.Child(html.SvgPath(html.AD("M16 12h.01"))),
		html.Child(html.SvgPath(html.AD("m17 7 5 5-5 5"))),
		html.Child(html.SvgPath(html.AD("m7 7-5 5 5 5"))),
		html.Child(html.SvgPath(html.AD("M8 12h.01"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
