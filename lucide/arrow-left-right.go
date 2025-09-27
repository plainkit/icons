package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowLeftRight creates a Arrow Left Right Lucide icon.
func ArrowLeftRight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-left-right", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M8 3 4 7l4 4"))),
		html.Child(html.SvgPath(html.AD("M4 7h16"))),
		html.Child(html.SvgPath(html.AD("m16 21 4-4-4-4"))),
		html.Child(html.SvgPath(html.AD("M20 17H4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
